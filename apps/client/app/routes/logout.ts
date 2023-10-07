import { redirect, type DataFunctionArgs } from "@remix-run/node";

import { destroySession, getSession } from "~/utils/session.server";

export const action = async ({ request }: DataFunctionArgs) => {
  const session = await getSession(request.headers.get("Cookie"));

  return redirect("/posts", {
    headers: {
      "Set-Cookie": await destroySession(session),
    },
  });
};
