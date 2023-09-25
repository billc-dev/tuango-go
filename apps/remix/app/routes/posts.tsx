import { useLoaderData } from "@remix-run/react";
import { useQuery } from "@tanstack/react-query";

export async function loader() {
  // const { data, error } = await serverClient.GET("/api/admin/v1/posts", {
  //   headers: {
  //     Authorization: "Bearer asdfasdf",
  //   },
  // });
  // if (error) {
  //   throw new Error(error.message);
  // }
  // return { postsData: data.data };
}

export default function Route() {
  const { postsData } = useLoaderData<typeof loader>();
  //   const revalidator = useRevalidator();

  const query = useQuery({
    // eslint-disable-next-line @tanstack/query/exhaustive-deps
    queryKey: ["posts"],
    initialData: postsData,
    queryFn: async () => {
      // const { data, error } = await client.GET("/api/admin/v1/posts", {});
      // if (error) {
      //   throw new Error(error.message);
      // }
      // return data.data;
    },
    // refetchInterval: 1000 * 5,
    gcTime: 1000,
    staleTime: 1000,
  });

  return (
    <div>
      {/* <Button onClick={() => revalidator.revalidate()}>Revalidate</Button> */}
      {/* {query.data?.posts?.map((post) => (
        <div key={post.id}>
          #{post.post_num} {post.title} {post.seller.display_name}
        </div>
      ))} */}
    </div>
  );
}
