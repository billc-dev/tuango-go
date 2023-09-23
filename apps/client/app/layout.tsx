import "@mantine/core/styles.css";
import "@mantine/dates/styles.css";
import "dayjs/locale/zh-tw";
import "./globals.css";

import type { Metadata } from "next";
import { ColorSchemeScript, MantineProvider } from "@mantine/core";

import Header from "../components/Header";

export const metadata: Metadata = {
  title: "開心團購",
  description: "",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html>
      <head>
        <ColorSchemeScript defaultColorScheme="auto" />
      </head>
      <body className="mb-2 h-screen overflow-hidden px-4 py-2 md:px-6 md:py-4">
        <MantineProvider
          defaultColorScheme="auto"
          theme={{
            primaryColor: "lime",
          }}
        >
          <Header />
          <div className="max-h-[100vh-60px] overflow-y-scroll">{children}</div>
        </MantineProvider>
      </body>
    </html>
  );
}
