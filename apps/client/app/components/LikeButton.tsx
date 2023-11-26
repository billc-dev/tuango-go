import { HeartIcon } from "@heroicons/react/24/outline";
import clsx from "clsx";
import { nanoid } from "nanoid";

import { useLiked, useLikePost, useUnlikePost } from "~/queries/likes";
import { useUser } from "~/queries/user";
import { useEnv } from "~/root";
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
  const user = useUser();
  const { LINE_CALLBACK_URL } = useEnv();

  function handleLike() {
    if (!user.data?.data) {
      return window.open(
        `https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=1654947889&redirect_uri=${LINE_CALLBACK_URL}&state=${nanoid()}&scope=profile%20openid`,
        "_self",
      );
      // return window.open(LINE_LOGIN_URL_WITH_PARAMS(`?redirect=${window.location.href}`), "_self");
    }
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
    <button type="button" onClick={() => handleLike()}>
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
