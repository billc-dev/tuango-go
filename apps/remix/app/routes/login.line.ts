import type { LoaderFunctionArgs } from "@remix-run/node";

export async function loader(req: LoaderFunctionArgs) {
  const url = new URL(req.request.url);
  const code = url.searchParams.get("code");
  //   const redirect = url.searchParams.get("redirect");

  const response = await fetch(
    `http://localhost:5010/api/client/v1/user/login/line/${code}?redirect_uri=${url.origin}/login/line`,
    { method: "POST" },
  );

  //   return { code };
  return await response.json();
}
