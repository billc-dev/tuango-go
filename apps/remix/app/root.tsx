import "@mantine/core/styles.css";
import "@mantine/dates/styles.css";

import { useState } from "react";
import { ColorSchemeScript, MantineProvider } from "@mantine/core";
import { cssBundleHref } from "@remix-run/css-bundle";
import type { LinksFunction } from "@remix-run/node";
import { Links, LiveReload, Meta, Outlet, Scripts, ScrollRestoration } from "@remix-run/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import stylesheet from "~/tailwind.css";
import Header from "./components/Header";

export const links: LinksFunction = () => [
  { rel: "stylesheet", href: stylesheet },
  ...(cssBundleHref ? [{ rel: "stylesheet", href: cssBundleHref }] : []),
];

export default function App() {
  const [queryClient] = useState(() => new QueryClient());

  return (
    <html lang="en" className="h-full">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width,initial-scale=1" />
        <Meta />
        <Links />
        <ColorSchemeScript defaultColorScheme="auto" />
      </head>
      <body className="mb-2 h-full px-4 py-2 md:px-6 md:py-4">
        <QueryClientProvider client={queryClient}>
          <MantineProvider
            defaultColorScheme="auto"
            theme={{
              primaryColor: "lime",
            }}
          >
            <Header />
            <div className="max-h-[100vh-60px] overflow-y-scroll">
              <Outlet />
            </div>
            <ScrollRestoration />
            <Scripts />
            <LiveReload />
          </MantineProvider>
        </QueryClientProvider>
      </body>
    </html>
  );
}
