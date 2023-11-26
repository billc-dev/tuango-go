import { useState } from "react";
import { MinusIcon, PlusIcon } from "@heroicons/react/24/outline";
import { useQueryClient } from "@tanstack/react-query";
import toast from "react-hot-toast";
import type { components } from "types/schema";

import { useScrollIntoView } from "~/hooks/useScrollIntoView";
import { useCreateOrder } from "~/mutations/orders";
import Card from "./Card";
import CardSubmitButton from "./CardSubmitButton";
import IconButton from "./IconButton";
import TextArea from "./TextArea";

interface PostOrderFormInterface {
  post: NonNullable<components["schemas"]["utils.Result-client_normalPost"]["data"]>;
  orders: components["schemas"]["utils.Result-array_client_postOrder"]["data"];
  scroll: boolean;
}

const PostOrderForm: React.FC<PostOrderFormInterface> = ({ post, orders, scroll }) => {
  const { ref } = useScrollIntoView(scroll);
  const queryClient = useQueryClient();

  const getInitialOrderItems = (post: PostOrderFormInterface["post"]) => {
    const sumOrders =
      orders?.reduce((sumOrders: Record<string, number>, orders) => {
        orders.order_items?.forEach((item) => {
          if (!item.post_item_id || !item.qty) return;
          sumOrders[item.post_item_id] = (sumOrders[item.post_item_id] ?? 0) + item.qty;
        });
        return sumOrders;
      }, {}) ?? {};

    return (
      post?.post_items?.map((item) => ({
        id: item.id ?? "",
        identifier: item.identifier,
        name: item.name,
        price: item.price ?? 0,
        stock: item.stock,
        qty: 0,
        orders: sumOrders[item.id!] ?? 0,
      })) ?? []
    );
  };

  const [orderItems, setOrderItems] = useState(getInitialOrderItems(post));
  const [comment, setComment] = useState("");

  const handleIncrementItemQty = ({ id, amount }: { id: string | undefined; amount: number }) => {
    if (!id) {
      return;
    }
    setOrderItems((orderForm) =>
      orderForm.map((item) => (item.id === id ? { ...item, qty: item.qty + amount } : item)),
    );
  };

  const sum = orderItems.reduce((sum, item) => sum + item.price * item.qty, 0);

  const createOrder = useCreateOrder();
  //   const orderCountQuery = useDeliveredOrderCount();

  const handleCreateOrder = async () => {
    if (!post.id) {
      return;
    }

    const order = orderItems.reduce((order: Record<string, number>, item) => {
      order[item.id] = item.qty;
      return order;
    }, {});

    // if (
    //   user.data &&
    //   orderCountQuery.data &&
    //   orderCountQuery.data.orderCount >= user.data.data.user.deliveredOrderCountLimit
    // ) {
    //   //   toast.error(
    //   //     `抱歉！您還有${orderCountQuery.data.orderCount}樣商品尚未領取，請儘速撥空領回，領取後系統將會繼續開放訂購！謝謝！`,
    //   //   );
    //   return;
    // }
    // const validatedOrderForm = await validateOrder(orderForm);
    // if (!validatedOrderForm) return;
    toast.loading("訂單製作中...", { id: "createOrder" });

    createOrder.mutate(
      { postId: post.id, order, comment, sum },
      {
        onSuccess: async (data, { postId }) => {
          toast.success("您的訂單已成立！", { id: "createOrder" });
          if (!data?.data?.post) {
            return;
          }
          setOrderItems(getInitialOrderItems(data.data.post));
          queryClient.setQueryData(["post", postId], data.data.post);
          queryClient.invalidateQueries({ queryKey: ["postOrders", postId] });
          // queryClient.invalidateQueries({ queryKey: ["posts"] });
        },
        onError: () => {
          toast.error("訂單製作失敗！", { id: "createOrder" });
        },
      },
    );
  };

  return (
    <Card>
      <div className="relative">
        <div ref={ref} className="absolute -top-12" />
      </div>
      <div className="flex flex-col items-center justify-center p-3">
        {orderItems.map((item) => {
          if (!item.id) {
            return undefined;
          }

          return (
            <div
              key={item.id}
              className="flex min-w-full flex-col items-center justify-center border-b border-zinc-300 py-2 last:border-b-0"
            >
              <div>
                <span>{`${item.identifier}.${item.name} $${item.price}`}</span>
                {item.orders && <span>~訂購數量{item.orders}</span>}
              </div>
              <div className="flex w-2/5 items-center justify-between">
                <IconButton
                  disabled={item.qty <= 0}
                  onClick={() => handleIncrementItemQty({ id: item.id, amount: -1 })}
                >
                  <MinusIcon />
                </IconButton>
                <div className="select-none px-2 text-xl">{item.qty}</div>
                <IconButton
                  disabled={item.qty >= (item.stock ?? 0)}
                  onClick={() => handleIncrementItemQty({ id: item.id, amount: 1 })}
                >
                  <PlusIcon />
                </IconButton>
              </div>
              {item.stock === 0 ? (
                <p className="text-red-500">已售完</p>
              ) : (
                <p>還剩{item.stock}件</p>
              )}
            </div>
          );
        })}
      </div>
      <div className="px-2 pb-2">
        <TextArea
          placeholder="備註"
          hiddenLabel
          maxRows={5}
          value={comment}
          onChange={(e) => setComment(e.target.value)}
        />
      </div>
      <CardSubmitButton
        disabled={sum <= 0}
        loading={createOrder.isPending}
        onClick={handleCreateOrder}
      >
        合計$ {sum}
      </CardSubmitButton>
    </Card>
  );
};

export default PostOrderForm;
