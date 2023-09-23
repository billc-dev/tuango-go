"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { Button, Group } from "@mantine/core";

const links = [
  { label: "取貨", link: "/" },
  { label: "進貨", link: "/deliver" },
  { label: "訊息", link: "/notifications" },
  { label: "貼文管理", link: "/posts" },
  { label: "訂單管理", link: "/orders" },
];

export default function Header() {
  const pathname = usePathname();

  return (
    <header className="mb-4">
      <Group gap={5}>
        {links.map((link) => (
          <Button
            key={link.label}
            component={Link}
            href={link.link}
            variant={pathname === link.link ? "filled" : "transparent"}
            className="py-0"
          >
            {link.label}
          </Button>
        ))}
      </Group>
    </header>
  );
}
