import { useContext, useRef, useState } from "react";
import { ArrowLeftOnRectangleIcon } from "@heroicons/react/24/outline";
import { Avatar, Button, Menu, Text } from "@mantine/core";
import { Form } from "@remix-run/react";
import { useQuery } from "@tanstack/react-query";
import { nanoid } from "nanoid";

import { UserContext } from "~/root";
import { client } from "~/utils/api";

// const links = [
//   { label: "取貨", link: "/" },
//   { label: "進貨", link: "/deliver" },
//   { label: "訊息", link: "/notifications" },
//   { label: "貼文管理", link: "/posts" },
//   { label: "訂單管理", link: "/orders" },
// ];

interface HeaderInterface {
  children?: React.ReactNode;
  LINE_CALLBACK_URL: string;
}

const Header: React.FC<HeaderInterface> = ({ children, LINE_CALLBACK_URL }) => {
  const ref = useRef<HTMLFormElement>(null);
  const [retries, setRetries] = useState(0);
  const { authenticated } = useContext(UserContext);
  const query = useQuery({
    enabled: authenticated,
    queryKey: ["user"],
    queryFn: async () => {
      const { data, error } = await client.GET("/api/client/v1/user", {});
      setRetries((retries) => retries + 1);
      if (error) {
        if (retries > 2) {
          ref.current?.submit();
        }
        throw new Error(error);
      }
      return { ...data };
    },
  });

  return (
    <header className="border-b border-zinc-200 px-2 py-2 shadow-sm dark:border-zinc-700 dark:bg-zinc-800">
      <div className="mx-auto flex max-w-4xl items-center justify-between">
        {query.data?.data ? (
          <div className="flex items-center justify-center gap-2">
            <Menu shadow="md">
              <Menu.Target>
                <Avatar src={query.data.data.picture_url} alt={query.data.data.display_name} />
              </Menu.Target>
              <Menu.Dropdown>
                <Form ref={ref} method="POST" action="/logout" reloadDocument>
                  <Menu.Item
                    type="submit"
                    leftSection={<ArrowLeftOnRectangleIcon className="h-6 w-6" />}
                  >
                    登出
                  </Menu.Item>
                </Form>
              </Menu.Dropdown>
            </Menu>
            <Text>會員編號: {query.data.data.pickup_num}</Text>
          </div>
        ) : !authenticated ? (
          <Button
            component="a"
            href={`https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=1654947889&redirect_uri=${LINE_CALLBACK_URL}&state=${nanoid()}&scope=profile%20openid`}
            className="!text-xl"
          >
            登入
          </Button>
        ) : (
          <div />
        )}
        <Form ref={ref} method="POST" action="/logout" reloadDocument></Form>

        {/* <Group gap={2}>
        {links.map((link) => (
          <Button
            key={link.label}
            component={Link}
            to={link.link}
            variant={location.pathname === link.link ? "outline" : "transparent"}
          >
            {link.label}
          </Button>
        ))}
      </Group> */}
        {children}
      </div>
    </header>
  );
};

export default Header;
