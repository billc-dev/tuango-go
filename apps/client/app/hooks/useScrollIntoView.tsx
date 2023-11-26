import { useEffect, useRef, useState } from "react";

export const useScrollIntoView = (enabled: boolean) => {
  const ref = useRef<HTMLDivElement>(null);
  const [scrolled, setScrolled] = useState(false);

  useEffect(() => {
    if (!ref?.current || !document || scrolled) return;

    if (enabled) {
      ref.current?.scrollIntoView({ behavior: "smooth" });
      setScrolled(true);
    }
  }, [scrolled, enabled, ref]);

  return { ref };
};
