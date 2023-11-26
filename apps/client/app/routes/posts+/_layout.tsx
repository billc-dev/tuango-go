import { Fragment } from "react";
import type { LoaderFunctionArgs } from "@remix-run/node";
import { json, redirect } from "@remix-run/node";
import { Outlet, type ShouldRevalidateFunction } from "@remix-run/react";
import { dehydrate, QueryClient, useInfiniteQuery } from "@tanstack/react-query";
import InfiniteScroll from "react-infinite-scroll-component";

import { PostCard } from "~/components/PostCard";
import { client, serverClient } from "~/utils/api";

export const shouldRevalidate: ShouldRevalidateFunction = () => {
  return false;
};

export const loader = async ({ request, params }: LoaderFunctionArgs) => {
  const { searchParams } = new URL(request.url);

  const legacyPostId = searchParams.get("postId"); // handle legacy postId
  if (legacyPostId) {
    return redirect(`/posts/${legacyPostId}`);
  }

  const queryClient = new QueryClient();

  const { postId } = params;

  if (!postId) {
    await queryClient.prefetchInfiniteQuery({
      queryKey: ["posts"],
      queryFn: async ({ pageParam }) => {
        const { data, error } = await serverClient.GET("/api/client/v1/posts", {
          params: {
            query: {
              page: pageParam,
            },
          },
        });
        if (error) {
          throw new Error(error);
        }
        return { ...data };
      },
      initialPageParam: 0,
    });
  }

  return json({ dehydratedState: dehydrate(queryClient) });
};

export default function Route() {
  const query = useInfiniteQuery({
    initialPageParam: 0,
    queryKey: ["posts"],
    queryFn: async ({ pageParam }) => {
      const { data, error } = await client.GET("/api/client/v1/posts", {
        params: {
          query: {
            page: pageParam,
          },
        },
      });
      if (error) {
        throw new Error(error);
      }
      return { ...data };
    },
    getNextPageParam: (_, pages) => pages.length,
  });

  return (
    <InfiniteScroll
      scrollableTarget="infiniteScrollTarget"
      className="flex w-full select-none flex-col items-center pb-16"
      loader={<div>loading</div>}
      next={() => query.fetchNextPage()}
      hasMore={query.hasNextPage}
      dataLength={query.data?.pages.reduce((sum, page) => (page.data?.length ?? 0) + sum, 0) || 0}
    >
      <div className="mb-2 grid grid-cols-2 gap-2 p-2 pb-0 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
        {query.data?.pages.map((page, index) => (
          <Fragment key={index}>
            {page.data?.map((post) => <PostCard key={post.id} post={post} />)}
          </Fragment>
        ))}
      </div>
      <Outlet />
      {/* <Modal /> */}
    </InfiniteScroll>
  );
}
