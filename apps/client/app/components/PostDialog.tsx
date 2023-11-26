import { useSearchParams } from "@remix-run/react";

import { usePost } from "~/queries/posts";
import { Dialog } from "./Dialog";
import { Post } from "./Post";

interface PostDialogInterface {
  postId: string;
}

export const PostDialog: React.FunctionComponent<PostDialogInterface> = ({ postId }) => {
  const postQuery = usePost(postId);
  const [searchParams] = useSearchParams();

  const action = searchParams.get("action");
  return (
    <Dialog title={postQuery.data?.title}>
      <div className="p-4">
        <Post postId={postId} action={action} />
      </div>
    </Dialog>
  );
};
