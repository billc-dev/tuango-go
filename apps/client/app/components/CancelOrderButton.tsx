import type { FC } from "react";
import { useState } from "react";
import { XCircleIcon } from "@heroicons/react/24/outline";
import type { components } from "types/schema";

import { useCancelOrder } from "~/mutations/orders";
import { usePost } from "~/queries/posts";
import { useUser } from "~/queries/user";
import Button from "./Button";
import IconButton from "./IconButton";
import NormalDialog from "./NormalDialog";

interface Props {
  order: components["schemas"]["client.postOrder"];
}

const CancelOrderButton: FC<Props> = ({ order }) => {
  const [open, setOpen] = useState(false);
  const userQuery = useUser();
  const postQuery = usePost(order.post_id);
  const cancelOrder = useCancelOrder();

  if (!userQuery.data || !postQuery.data || !order.id || !order.post_id) {
    return null;
  }

  if (
    order.user_id !== userQuery.data.data?.id ||
    order.status !== "ordered" ||
    postQuery.data.status !== "open"
  ) {
    return null;
  }

  return (
    <IconButton disabled={cancelOrder.isPending} onClick={() => setOpen(true)}>
      <XCircleIcon className="text-zinc-500" />
      <NormalDialog open={open} setOpen={setOpen} title="您確定要取消這筆訂單嗎？">
        <div className="flex justify-end gap-2 pt-2">
          <Button
            size="lg"
            variant="danger"
            loading={cancelOrder.isPending}
            onClick={() =>
              cancelOrder.mutate(
                { orderId: order.id!, postId: order.post_id! },
                { onSuccess: () => setOpen(false) },
              )
            }
          >
            確定
          </Button>
          <Button size="lg" onClick={() => setOpen(false)}>
            取消
          </Button>
        </div>
      </NormalDialog>
    </IconButton>
  );
};

export default CancelOrderButton;
