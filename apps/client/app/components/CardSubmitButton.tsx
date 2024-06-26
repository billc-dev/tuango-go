import type { ButtonHTMLAttributes, FC } from "react";

import AnimatedSpinner from "./AnimatedSpinner";

interface Props extends ButtonHTMLAttributes<HTMLButtonElement> {
  loading?: boolean;
}

const CardSubmitButton: FC<Props> = ({ children, loading, disabled, ...props }) => {
  return (
    <button
      disabled={disabled || loading}
      className="bg-line-400 hover:bg-line-700 active:bg-line-800 flex h-12 w-full items-center justify-center text-lg text-white transition disabled:bg-zinc-300 dark:disabled:bg-zinc-600"
      {...props}
    >
      {loading ? <AnimatedSpinner /> : children}
    </button>
  );
};

export default CardSubmitButton;
