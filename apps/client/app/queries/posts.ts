import { useQuery } from "@tanstack/react-query";

import { client } from "~/utils/api";
import { useUser } from "./user";

export const usePost = (postId: string | null | undefined) => {
  return useQuery({
    enabled: Boolean(postId),
    queryKey: ["post", postId],
    queryFn: async () => {
      const { data, error } = await client.GET("/api/client/v1/posts/{id}", {
        params: {
          path: {
            id: postId ?? "",
          },
        },
      });
      if (error) {
        throw new Error(error);
      }
      if (!data?.data) {
        throw new Error("No data");
      }
      return data.data;
    },
  });
};

export const usePostOrders = (postId: string | undefined) => {
  const userQuery = useUser();

  const postOrdersQuery = useQuery({
    enabled: Boolean(userQuery.data?.data),
    queryKey: ["postOrders", postId],
    queryFn: async () => {
      const { data, error } = await client.GET("/api/client/v1/posts/{id}/orders", {
        params: {
          path: {
            id: postId ?? "",
          },
        },
      });
      if (error) {
        throw new Error(error);
      }
      if (!data?.data) {
        throw new Error("No data");
      }

      return data.data;
    },
  });

  return { userQuery, postOrdersQuery };
};
