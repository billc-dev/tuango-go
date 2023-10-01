import type { LoaderFunctionArgs } from "@remix-run/node";

import { serverClient } from "~/utils/api";

export async function loader(req: LoaderFunctionArgs) {
  console.log("API request:", `/api/${req.params["*"]}`);

  const url = new URL(req.request.url);
  const query = Object.fromEntries(new URLSearchParams(url.searchParams));
  console.log("query", query);

  const { data, error } = await serverClient.GET(
    // @ts-ignore
    `/api/${req.params["*"]}`,
    {
      headers: {
        Authorization: "Bearer asdfasdf",
      },
      params: {
        query,
      },
    },
  );
  if (error instanceof Error) {
    throw new Error(error.message);
  }
  return data;
}
