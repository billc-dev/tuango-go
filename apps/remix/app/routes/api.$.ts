import type { LoaderFunctionArgs } from "@remix-run/node";

export async function loader(req: LoaderFunctionArgs) {
  console.log("Remix proxy:", `/api/${req.params["*"]}`);

  // const { data, error } = await serverClient.GET(`/api/${req.params["*"]}`, {
  //   headers: {
  //     Authorization: "Bearer asdfasdf",
  //   },
  // });
  // if (error) {
  //   throw new Error(error.message);
  // }
  // return data;
}
