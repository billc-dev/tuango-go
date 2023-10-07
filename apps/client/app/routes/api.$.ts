import { type DataFunctionArgs } from "@remix-run/node";

import { serverClient } from "~/utils/api";
import { getSession } from "~/utils/session.server";

export async function loader({ request, params }: DataFunctionArgs) {
  const url = new URL(request.url);
  const query = Object.fromEntries(new URLSearchParams(url.searchParams));

  console.log("API request:", `/api/${params["*"]}`);
  console.log("query", query);

  const session = await getSession(request.headers.get("Cookie"));
  const token = session.get("token");

  const { data, error } = await serverClient.GET(
    // @ts-ignore
    `/api/${params["*"]}`,
    {
      headers: {
        Authorization: token ? `Bearer ${token}` : "",
      },
      params: {
        query,
      },
    },
  );

  if (error) {
    console.error(error);
    throw new Error(error);
  }

  return data;
}

export async function action({ request, params }: DataFunctionArgs) {
  const url = new URL(request.url);
  const query = Object.fromEntries(new URLSearchParams(url.searchParams));

  console.log("POST API request:", `/api/${params["*"]}`);
  console.log("query", query);

  const session = await getSession(request.headers.get("Cookie"));
  const token = session.get("token");
  console.log("url", url.searchParams.toString());

  const response = await fetch(
    `http://127.0.0.1:5010/api/${params["*"]}?${url.searchParams.toString()}`,
    {
      method: request.method,
      headers: {
        Authorization: token ? `Bearer ${token}` : "",
      },
      body: request.body,
    },
  );

  return await response.json();
}
