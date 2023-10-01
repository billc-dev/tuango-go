import { Avatar, Text } from "@mantine/core";

// const links = [
//   { label: "取貨", link: "/" },
//   { label: "進貨", link: "/deliver" },
//   { label: "訊息", link: "/notifications" },
//   { label: "貼文管理", link: "/posts" },
//   { label: "訂單管理", link: "/orders" },
// ];

interface HeaderInterface {
  children?: React.ReactNode;
}

const Header: React.FC<HeaderInterface> = ({ children }) => {
  // const location = useLocation();
  return (
    <header className="flex items-center justify-between border-b border-zinc-200 px-2 py-2 shadow-sm dark:border-zinc-700 dark:bg-zinc-800">
      <div className="flex items-center justify-center gap-2">
        <Avatar alt="User">BC</Avatar>
        <Text>會員編號: 1589</Text>
      </div>

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
    </header>
  );
};

export default Header;
