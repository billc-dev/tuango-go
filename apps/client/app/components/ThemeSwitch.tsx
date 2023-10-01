import { MoonIcon, SunIcon } from "@heroicons/react/24/outline";
import { ActionIcon, useMantineColorScheme } from "@mantine/core";

export default function ThemeSwitch() {
  const { toggleColorScheme } = useMantineColorScheme();

  return (
    <ActionIcon variant="subtle" color="gray" size="lg" onClick={() => toggleColorScheme()}>
      <SunIcon className="hidden h-6 w-6 text-white dark:block" />
      <MoonIcon className="block h-6 w-6 text-black dark:hidden" />
    </ActionIcon>
  );
}
