import { usePost, usePostOrders } from "~/queries/posts";
import { getRelativeDate } from "~/utils/date";
import { getOrderStatusLabel } from "~/utils/text";
import CancelOrderButton from "./CancelOrderButton";
import CardHeader from "./CardHeader";
import PostOrderForm from "./PostOrderForm";

interface PostOrdersInterface {
  postId: string;
  scroll: boolean;
}

export const PostOrders: React.FC<PostOrdersInterface> = ({ postId, scroll }) => {
  const { userQuery, postOrdersQuery } = usePostOrders(postId);
  const postQuery = usePost(postId);
  const isSeller = userQuery.data?.data?.id === postQuery.data?.seller_id;
  const isAdmin = userQuery.data?.data?.role === "admin";

  return (
    <>
      {postOrdersQuery.data?.map((order) => (
        <div key={order.id} className="my-4 ml-0.5 rounded-lg text-sm first:mt-6">
          <CardHeader
            img={order.user?.picture_url}
            title={order.user?.display_name}
            titleDate={getRelativeDate(order.created_at)}
            subtitle={`序號: ${order.order_num} ${getOrderStatusLabel(order.status)}`}
            subtitleResetStyles
            action={
              <>
                <CancelOrderButton order={order} />
              </>
            }
          />
          <div className="ml-12">
            {order.order_items?.map((item) => (
              <ul key={item.id}>
                {item.identifier}. {item.name}+{item.qty}{" "}
                {(isSeller || isAdmin) && "$" + (item.qty ?? 0) * (item.price ?? 0)}
              </ul>
            ))}
            {order.comment && <p className="whitespace-pre pt-1 text-sm">備註: {order.comment}</p>}
            {/* {post?.status !== "open" && isSeller && (
        <HasNameButton {...{ postId: post._id, order }} />
      )} */}
          </div>
        </div>
      ))}
      {postQuery.data && postOrdersQuery.data && (
        <PostOrderForm post={postQuery.data} orders={postOrdersQuery.data} scroll={scroll} />
      )}
    </>
  );
};
