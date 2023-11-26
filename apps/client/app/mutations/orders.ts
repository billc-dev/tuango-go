import { useMutation, useQueryClient } from "@tanstack/react-query";
import toast from "react-hot-toast";

import { client } from "~/utils/api";

export const useCreateOrder = () => {
  return useMutation({
    mutationFn: async ({
      postId,
      order,
      comment,
      sum,
    }: {
      postId: string;
      sum: number;
      comment: string;
      order: Record<string, number>;
    }) => {
      const { data, error } = await client.POST("/api/client/v1/orders", {
        body: {
          postId,
          order,
          comment,
          sum,
        },
      });

      if (error) {
        throw new Error(error);
      }

      return data;
    },
  });
};

export const useCancelOrder = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async ({ postId, orderId }: { postId: string; orderId: string }) => {
      toast.loading("取消中...", { id: "cancelOrder" });
      const { error } = await client.DELETE("/api/client/v1/orders/{id}", {
        params: {
          path: {
            id: orderId,
          },
        },
      });

      if (error) {
        throw new Error(error);
      }

      const ordersData = queryClient.getQueryData(["postOrders", postId]);
      // TODO: fix
      queryClient.setQueryData(
        ["postOrders", postId],
        ordersData.filter((order) => order.id !== orderId),
      );
    },
    onError: () => {
      toast.error("取消失敗！", { id: "cancelOrder" });
    },
    onSuccess: () => {
      toast.success("您的訂單已取消！", { id: "cancelOrder" });
    },
  });
};
