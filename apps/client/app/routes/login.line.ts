import { redirect, type LoaderFunctionArgs } from "@remix-run/node";

import { serverClient } from "~/utils/api";
import { commitSession, getSession } from "~/utils/session.server";

export async function loader({ request }: LoaderFunctionArgs) {
  const url = new URL(request.url);
  const code = url.searchParams.get("code");
  //   const redirect = url.searchParams.get("redirect");

  if (!code) {
    throw new Error("Code is not present");
  }

  const { data, error } = await serverClient.POST("/api/client/v1/user/login/line/{code}", {
    params: {
      path: {
        code,
      },
      query: {
        redirect_uri: process.env.LINE_CALLBACK_URL,
      },
    },
  });

  if (error) {
    throw new Error(error);
  }

  if (!data?.data) {
    throw new Error("JWT token is not present");
  }

  const session = await getSession(request.headers.get("Cookie"));
  session.set("token", data.data);

  return redirect("/posts", {
    headers: {
      "Set-Cookie": await commitSession(session),
    },
  });
}
