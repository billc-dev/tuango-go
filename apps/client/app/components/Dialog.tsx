import { useEffect, useRef, useState } from "react";
import { ArrowLeftIcon } from "@heroicons/react/24/outline";
import { useNavigate } from "@remix-run/react";
import clsx from "clsx";

import IconButton from "./IconButton";

interface DialogInterface {
  title?: string;
  action?: React.ReactNode;
  children: React.ReactNode;
}

export const Dialog: React.FunctionComponent<DialogInterface> = ({ title, action, children }) => {
  const [animate, setAnimate] = useState(false);
  const navigate = useNavigate();
  const ref = useRef<HTMLDivElement>();
  useEffect(() => {
    // if (open) {
    // document.bodymoveProperty("overflow");
    // };.style.overflow = "hidden";
    setAnimate(true);
    ref.current?.focus();
    // }
    // return () => {
    //   document.body.style.re
  }, []);

  const handleClose = () => {
    navigate("..", { relative: "path" });
  };

  return (
    <>
      <dialog
        open
        className={clsx(
          "fixed inset-0 left-0 top-0 z-20 h-full w-full overflow-x-hidden overflow-y-scroll overscroll-y-contain bg-white p-0 transition-opacity duration-300 dark:bg-zinc-900 md:bottom-0 md:max-h-[95%] md:w-fit md:min-w-[512px] md:max-w-md md:rounded md:shadow-lg md:ring-zinc-700 md:dark:ring-1",
          animate ? "opacity-100" : "opacity-0",
        )}
      >
        <div className="sticky top-0 z-20 flex w-full items-center bg-white p-1 shadow dark:bg-zinc-800">
          <div className="mr-3 flex w-full items-center justify-between">
            <div className="flex items-center">
              <IconButton onClick={handleClose}>
                <ArrowLeftIcon />
              </IconButton>
              <h1 className="line-clamp-1 select-text pr-2 text-xl dark:text-white">{title}</h1>
            </div>
            <div>{action}</div>
          </div>
        </div>
        {children}
      </dialog>
      <div
        className={clsx(
          "fixed left-0 top-0 hidden h-screen w-screen bg-zinc-900/50 transition-opacity duration-300 dark:bg-zinc-900/60 md:block",
          animate ? "opacity-100" : "opacity-0",
        )}
        onClick={handleClose}
      />
    </>
  );
};
