import { useState } from "react";

import { useScrollIntoView } from "~/hooks/useScrollIntoView";
import Card from "./Card";
import CardSubmitButton from "./CardSubmitButton";
import TextArea from "./TextArea";

const Comment = () => {
  return (
    <>
      <CommentForm />
    </>
  );
};

export default Comment;

const CommentForm = () => {
  const isLoading = false;
  const { ref } = useScrollIntoView(isLoading);
  const [comment, setComment] = useState("");

  return (
    <Card>
      <div className="relative px-2 pt-3">
        <div ref={ref} className="absolute -top-12" />
        <TextArea
          autoFocus
          hiddenLabel
          placeholder="問題"
          value={comment}
          onChange={(e) => setComment(e.target.value)}
        />
      </div>
      <CardSubmitButton disabled={!comment} onClick={() => handleCreateComment()}>
        新增問題
      </CardSubmitButton>
    </Card>
  );
};
