import { useEffect, useRef } from "react";
import nProgress from "nprogress";

export const useProgress = (isLoading: boolean) => {
  let progressBarTimeout = useRef<NodeJS.Timeout>();

  useEffect(() => {
    const startProgressBar = () => {
      clearTimeout(progressBarTimeout.current);
      progressBarTimeout.current = setTimeout(nProgress.start, 200);
    };

    const stopProgressBar = () => {
      clearTimeout(progressBarTimeout.current);
      nProgress.done();
    };

    if (isLoading) {
      startProgressBar();
    } else {
      stopProgressBar();
    }
  }, [isLoading]);
};
