import "react-lazy-load-image-component/src/effects/opacity.css";
import "yet-another-react-lightbox/plugins/counter.css";
import "yet-another-react-lightbox/plugins/thumbnails.css";
import "yet-another-react-lightbox/styles.css";

import { useState } from "react";
import {
  ChatBubbleLeftEllipsisIcon,
  ChevronLeftIcon,
  ChevronRightIcon,
  ClipboardDocumentListIcon,
} from "@heroicons/react/24/outline";
import { Carousel } from "@mantine/carousel";
import { Avatar } from "@mantine/core";
import { useSearchParams } from "@remix-run/react";
import clsx from "clsx";
import { LazyLoadImage } from "react-lazy-load-image-component";
import Lightbox from "yet-another-react-lightbox";
import Counter from "yet-another-react-lightbox/plugins/counter";
import Thumbnails from "yet-another-react-lightbox/plugins/thumbnails";
import Zoom from "yet-another-react-lightbox/plugins/zoom";

import { usePost } from "~/queries/posts";
import { date, getFullDateFromNow } from "~/utils/date";
import { getStorageTypeLabel } from "~/utils/text";
import Comment from "./Comment";
import LikeButton from "./LikeButton";
import { PostOrders } from "./PostOrders";
import TabButton from "./TabButton";
import TabContainer from "./TabContainer";

interface PostInterface {
  postId: string;
  action: string | null;
}

export const Post: React.FC<PostInterface> = ({ postId, ...props }) => {
  const postQuery = usePost(postId);
  const [imageIndex, setImageIndex] = useState(0);
  const [openLightbox, setOpenLightbox] = useState(false);
  const [searchParams, setSearchParams] = useSearchParams();
  const [action, setAction] = useState(props.action);

  if (!postQuery.data) {
    return null;
  }
  const { data: post } = postQuery;

  return (
    <>
      <div className="flex">
        <Avatar src={post.seller?.picture_url} alt={post.seller?.display_name} />
        <div className="ml-1 flex flex-col pl-2">
          <div className="line-clamp-1 text-sm">{post.seller?.display_name}</div>
          <div className="line-clamp-1 text-xs text-zinc-400">
            {getFullDateFromNow(post.created_at)}
          </div>
        </div>
      </div>
      <Carousel
        className={clsx("relative -mx-4 mt-2 min-w-full")}
        onSlideChange={(index) => setImageIndex(index)}
        withIndicators
        speed={30}
        height={300}
        controlsOffset="xs"
        previousControlIcon={<ChevronLeftIcon className="h-8 w-8" />}
        nextControlIcon={<ChevronRightIcon className="h-8 w-8" />}
        classNames={{
          indicators: "absolute px-2 ",
          indicator: "!bg-zinc-600 dark:!bg-white",
          control:
            "!bg-transparent !outline-none !border-none text-black dark:text-white !shadow-none",
        }}
      >
        {post.images?.map((image, index) => (
          <Carousel.Slide key={index} className="flex items-center justify-center">
            <LazyLoadImage
              src={image.md}
              visibleByDefault={imageIndex + 1 === index}
              effect="opacity"
              placeholder={<div className="h-72 w-full bg-zinc-300 dark:bg-zinc-600" />}
              className={`max-h-72 object-contain transition-all duration-300`}
              onClick={() => setOpenLightbox(true)}
            />
          </Carousel.Slide>
        ))}
      </Carousel>
      {openLightbox && (
        <Lightbox
          open={openLightbox}
          close={() => setOpenLightbox(false)}
          slides={post.images?.map((image) => ({
            src: image.md ?? "",
          }))}
          index={imageIndex}
          plugins={[Counter, Thumbnails, Zoom]}
        />
      )}
      <div className="select-text py-4">
        <p className="font-bold">
          #{post.post_num} {post.title} #{post.seller?.display_name}
        </p>
        <p>❤️ #結單日: {post.deadline ? date(post.deadline) : "成團通知"}</p>
        <p>💚 #到貨日: {post.delivery_date ? date(post.delivery_date) : "貨到通知"}</p>
        <p>儲存方式: {getStorageTypeLabel(post.storage_type)}</p>
        <p className="whitespace-pre-line pt-4">{post.body?.trim()}</p>
      </div>
      <TabContainer>
        <LikeButton tabButton postId={postId} likeCount={post.like_count} />
        <TabButton
          icon={<ChatBubbleLeftEllipsisIcon />}
          selected={action === "comment"}
          onClick={() => setAction("comment")}
        >
          {post.comment_count} 問與答
        </TabButton>
        <TabButton
          icon={<ClipboardDocumentListIcon />}
          selected={action === "order" || action === null}
          onClick={() => {
            setAction("order");
            searchParams.set("action", "order");
            setSearchParams(searchParams, { preventScrollReset: true });
          }}
        >
          {post.status !== "completed" && post.order_count} 訂單
        </TabButton>
      </TabContainer>
      {action !== "comment" ? (
        <PostOrders postId={postId} scroll={action === "order"} />
      ) : (
        <Comment />
      )}
    </>
  );
};
