import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import type { components } from "types/schema";

import { client } from "~/utils/api";
import { useUser } from "./user";

export const useLikes = () => {
  const { data } = useUser();

  return useQuery({
    enabled: Boolean(data?.data),
    queryKey: ["likes"],
    queryFn: async () => {
      const { data, error } = await client.GET("/api/client/v1/user/likes", {});
      if (error) {
        throw new Error(error);
      }
      return { ...data };
    },
  });
};

export const useLiked = (postId?: string) => {
  const likesQuery = useLikes();
  return likesQuery.data?.data?.some((like) => like.post_id === postId);
};

export const useLikePost = (postId: string | undefined) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (postId: string) => {
      const { data, error } = await client.POST("/api/client/v1/posts/{id}/like", {
        params: {
          path: {
            id: postId,
          },
        },
      });

      if (error) {
        throw new Error(error);
      }

      return data;
    },
    onMutate: async () => {
      await queryClient.cancelQueries({ queryKey: ["likes"] });
      const previousLikes = queryClient.getQueryData(["likes"]);
      queryClient.setQueryData<components["schemas"]["utils.Result-array_ent_Like"]>(
        ["likes"],
        (old) => {
          return {
            data: old?.data ? [...old?.data, { post_id: postId }] : [{ post_id: postId }],
          };
        },
      );
      return { previousLikes };
    },
    onError: (_, __, context) => {
      queryClient.setQueryData(["likes"], context?.previousLikes);
    },
    onSettled: async () => {
      queryClient.invalidateQueries({ queryKey: ["likes"] });
      queryClient.invalidateQueries({ queryKey: ["post", postId] });
    },
  });
};

export const useUnlikePost = (postId: string | undefined) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (postId: string) => {
      const { data, error } = await client.DELETE("/api/client/v1/posts/{id}/like", {
        params: {
          path: {
            id: postId,
          },
        },
      });

      if (error) {
        throw new Error(error);
      }

      return data;
    },
    onMutate: async () => {
      await queryClient.cancelQueries({ queryKey: ["likes"] });
      const previousLikes = queryClient.getQueryData(["likes"]);
      queryClient.setQueryData<components["schemas"]["utils.Result-array_ent_Like"]>(
        ["likes"],
        (old) => {
          return {
            data: old?.data ? old.data.filter((like) => like.post_id !== postId) : [],
          };
        },
      );
      return { previousLikes };
    },
    onError: (_, __, context) => {
      queryClient.setQueryData(["likes"], context?.previousLikes);
    },
    onSettled: async () => {
      queryClient.invalidateQueries({ queryKey: ["likes"] });
      queryClient.invalidateQueries({ queryKey: ["post", postId] });
    },
  });
};
