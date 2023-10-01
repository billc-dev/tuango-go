import type { Config } from "tailwindcss";

export default {
  content: ["./app/**/*.{js,jsx,ts,tsx}"],
  darkMode: ["class", '[data-mantine-color-scheme="dark"]'],
  theme: {
    extend: {},
  },
  plugins: [],
} satisfies Config;
