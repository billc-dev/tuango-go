import {
  ChatBubbleOvalLeftEllipsisIcon,
  ClipboardDocumentListIcon,
  DocumentPlusIcon,
  MagnifyingGlassIcon,
  ShoppingCartIcon,
} from "@heroicons/react/24/outline";
import { Link, useLocation } from "@remix-run/react";
import clsx from "clsx";

const links = [
  { label: "開心團購", link: "/posts", icon: <img alt="logo" src="/logo.png" /> },
  { label: "現貨區", link: "/super-buy", icon: <ShoppingCartIcon /> },
  { label: "新增貼文", link: "/create-post", icon: <DocumentPlusIcon /> },
  { label: "聊天室", link: "/chat", icon: <ChatBubbleOvalLeftEllipsisIcon /> },
  { label: "我的訂單", link: "/orders", icon: <ClipboardDocumentListIcon /> },
  { label: "搜尋", link: "/search", icon: <MagnifyingGlassIcon /> },
];

export default function BottomNavbar() {
  const location = useLocation();
  return (
    <div className="z-10 flex gap-2 rounded-t-2xl bg-white ring-1 ring-zinc-300 dark:bg-zinc-800 dark:ring-zinc-700">
      {links.map((link) => (
        <Link
          key={link.link}
          to={link.link}
          prefetch="intent"
          className={`$ w-full transform rounded-lg py-2 transition active:bg-zinc-200 dark:active:bg-zinc-700`}
        >
          <div
            className={clsx(
              "flex flex-col items-center transition",
              location.pathname === link.link && "scale-[1.05] text-blue-600 dark:text-blue-400",
            )}
          >
            <div className="h-6 w-6">{link.icon}</div>
            <label className="select-none text-xs">{link.label}</label>
          </div>
        </Link>
      ))}
    </div>
  );
}
