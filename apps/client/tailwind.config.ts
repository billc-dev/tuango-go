import type { Config } from "tailwindcss";

export default {
  content: ["./app/**/*.{js,jsx,ts,tsx}"],
  darkMode: ["class", '[data-mantine-color-scheme="dark"]'],
  theme: {
    extend: {
      colors: {
        line: {
          300: "#00CC00",
          400: "#00B900",
          500: "#00A800",
          600: "#00B900",
          700: "#009900",
          800: "#008B00",
          900: "#007E00",
        },
      },
    },
  },
  plugins: [],
} satisfies Config;
