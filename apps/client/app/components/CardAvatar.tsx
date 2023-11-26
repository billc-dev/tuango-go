import { useState } from "react";
import { LazyLoadImage } from "react-lazy-load-image-component";

interface Props {
  src: string | undefined;
  alt: string;
  notifications?: number;
}

const CardAvatar: React.FunctionComponent<Props> = ({ src, alt, notifications }) => {
  const [error, setError] = useState(false);
  const [loaded, setLoaded] = useState(false);
  return (
    <div className="relative min-w-[40px] select-none">
      {!error || !src ? (
        <LazyLoadImage
          alt={alt}
          className={`h-10 w-10 overflow-hidden rounded-full transition-all duration-300 ${
            loaded ? "opacity-100 blur-0" : "opacity-0 blur-sm"
          }`}
          src={src}
          onLoad={() => setLoaded(true)}
          placeholder={<div className="h-10 w-10 rounded-full bg-zinc-400" />}
          onError={() => setError(true)}
        />
      ) : (
        <div className="flex h-10 w-10 items-center justify-center rounded-full bg-zinc-400 text-xl text-white dark:bg-zinc-500">
          {alt.substring(0, 1)}
        </div>
      )}
      {notifications ? (
        <div className="absolute -right-2.5 -top-2.5 rounded-full bg-red-600 text-sm text-white">
          <div className="mx-2 my-0.5">{notifications}</div>
        </div>
      ) : null}
    </div>
  );
};

export default CardAvatar;
