import type { LoaderFunctionArgs } from "@remix-run/node";
import { json } from "@remix-run/node";
import { useLoaderData, type ShouldRevalidateFunction } from "@remix-run/react";
import { dehydrate, QueryClient } from "@tanstack/react-query";

import { PostDialog } from "~/components/PostDialog";
import { serverClient } from "~/utils/api";

export const shouldRevalidate: ShouldRevalidateFunction = () => {
  return false;
};

export const loader = async ({ params }: LoaderFunctionArgs) => {
  const queryClient = new QueryClient();

  const { postId } = params;

  if (postId) {
    await queryClient.prefetchQuery({
      queryKey: ["post", postId],
      queryFn: async () => {
        const { data, error } = await serverClient.GET("/api/client/v1/posts/{id}", {
          params: {
            path: {
              id: postId,
            },
          },
        });
        if (error) {
          throw new Error(error);
        }
        return data?.data;
      },
    });
  }

  return json({ dehydratedState: dehydrate(queryClient), postId: params.postId });
};

export default function Route() {
  const { postId } = useLoaderData<typeof loader>();
  return postId ? <PostDialog postId={postId} /> : null;
}
