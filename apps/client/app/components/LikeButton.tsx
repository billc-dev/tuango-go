import { HeartIcon } from "@heroicons/react/24/outline";
import clsx from "clsx";

import { useLiked, useLikePost, useUnlikePost } from "~/queries/likes";
import TabButton from "./TabButton";

interface LikeButtonInterface {
  postId: string;
  likeCount?: number;
  tabButton?: boolean;
}

const LikeButton: React.FC<LikeButtonInterface> = ({ postId, tabButton, likeCount }) => {
  const liked = useLiked(postId);
  const likePost = useLikePost(postId);
  const unlikePost = useUnlikePost(postId);

  function handleLike() {
    // if (!userQuery.data?.data.user) {
    //   return window.open(LINE_LOGIN_URL_WITH_PARAMS(`?redirect=${window.location.href}`), "_self");
    // }
    if (!liked) {
      return likePost.mutate(postId);
    }
    return unlikePost.mutate(postId);
  }
  return tabButton ? (
    <TabButton selected={false} onClick={() => handleLike()}>
      <div className="mr-1">
        <HeartIcon
          className={clsx(
            "h-5 w-5  stroke-red-600 transition-colors",
            liked ? "fill-red-600" : "fill-transparent",
          )}
        />
      </div>
      {likeCount} 喜歡
    </TabButton>
  ) : (
    <button
      type="button"
      onClick={() => {
        if (!postId) {
          return;
        }
        !liked ? likePost.mutate(postId) : unlikePost.mutate(postId);
      }}
    >
      <HeartIcon
        className={clsx(
          "h-6 w-6  stroke-red-600 transition-colors",
          liked ? "fill-red-600" : "fill-transparent",
        )}
      />
    </button>
  );
};

export default LikeButton;
