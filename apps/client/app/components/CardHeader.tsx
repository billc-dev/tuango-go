import clsx from "clsx";

import CardAvatar from "./CardAvatar";

interface Props {
  img: string | undefined;
  title: string | undefined;
  titleDate?: string;
  subtitle?: string;
  subtitleResetStyles?: boolean;
  action?: JSX.Element | null;
  notifications?: number;
  onClick?: () => void;
}

const CardHeader: React.FC<Props> = ({
  img,
  title = "",
  titleDate,
  subtitle,
  subtitleResetStyles,
  action,
  notifications,
  onClick,
}) => {
  return (
    <div className="flex items-center justify-between">
      <div className="flex">
        <CardAvatar
          src={img !== "undefined/small" ? img : undefined}
          alt={title}
          notifications={notifications}
        />
        <div className="ml-1 flex flex-col pl-1">
          <div className="flex items-center">
            <p className="line-clamp-1 text-sm">{title}</p>
            {titleDate && (
              <p className="line-clamp-1 min-w-[60px] pl-1 text-xs text-zinc-400">{titleDate}</p>
            )}
          </div>
          <p
            className={clsx(
              "line-clamp-1",
              !subtitleResetStyles ? "text-xs text-zinc-400" : "text-sm",
            )}
          >
            {subtitle}
          </p>
        </div>
      </div>
      <div>{action}</div>
    </div>
  );
};

export default CardHeader;
