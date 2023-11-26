import React from "react";
import type { TextareaAutosizeProps } from "react-textarea-autosize";
import TextareaAutosize from "react-textarea-autosize";

interface Props extends TextareaAutosizeProps, React.RefAttributes<HTMLTextAreaElement> {
  color?: "grey";
  label?: string;
  error?: string;
  hiddenLabel?: boolean;
}

const TextArea: React.FC<Props> = React.forwardRef((props, ref) => {
  const { color, hiddenLabel, error, ...rest } = props;
  return (
    <>
      {!hiddenLabel
        ? props.placeholder && (
            <label htmlFor={props.placeholder} className="block pb-2">
              {props.placeholder}
            </label>
          )
        : null}
      <TextareaAutosize
        className={`mb-2 w-full rounded-lg border border-zinc-200 px-3 py-4 placeholder-zinc-400 dark:border-zinc-600 dark:bg-zinc-800 ${
          color === "grey" && "bg-zinc-100"
        } ${
          !error
            ? "focus:border-line-400 focus:ring-line-400 focus:ring-1"
            : "border-red-500 ring-1 ring-red-500"
        }`}
        ref={ref}
        {...rest}
      />
      {error && <p className="-mt-1 text-center text-sm text-red-600">{error}</p>}
    </>
  );
});

TextArea.displayName = "TextArea";

export default TextArea;
