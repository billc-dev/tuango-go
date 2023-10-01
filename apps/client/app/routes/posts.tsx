import { Fragment } from "react";
import type { LoaderFunctionArgs } from "@remix-run/node";
import { json } from "@remix-run/node";
import type { ShouldRevalidateFunction } from "@remix-run/react";
import { dehydrate, QueryClient, useInfiniteQuery } from "@tanstack/react-query";
import InfiniteScroll from "react-infinite-scroll-component";

import { Modal } from "~/components/Modal";
import { PostCard } from "~/components/PostCard";
import { client } from "~/utils/api";

export const shouldRevalidate: ShouldRevalidateFunction = () => {
  return false;
};

export const loader = async (req: LoaderFunctionArgs) => {
  const queryClient = new QueryClient();

  const post_id = new URL(req.request.url).searchParams.get("post_id");

  if (post_id) {
    await queryClient.prefetchQuery({
      queryKey: ["post", post_id],
      queryFn: async () => {
        const { data, error } = await client.GET("/api/client/v1/posts/{id}", {
          params: {
            path: {
              id: post_id,
            },
          },
        });
        if (error) {
          throw new Error(error.message);
        }
        return { ...data };
      },
    });
  } else {
    await queryClient.prefetchInfiniteQuery({
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
          throw new Error(error.message);
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
        throw new Error(error.message);
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
      {/* <Button onClick={open}>Open</Button> */}
      <div className="mb-2 grid grid-cols-2 gap-2 p-2 pb-0 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
        {query.data?.pages.map((page, index) => (
          <Fragment key={index}>
            {page.data?.map((post) => <PostCard key={post.id} post={post} />)}
          </Fragment>
        ))}
      </div>
      <Modal />
    </InfiniteScroll>
  );
}
