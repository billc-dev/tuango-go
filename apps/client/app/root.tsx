import "@mantine/core/styles.css";
import "@mantine/dates/styles.css";
import "@mantine/carousel/styles.css";

import { useEffect, useMemo, useRef, useState } from "react";
import { ColorSchemeScript, MantineProvider } from "@mantine/core";
import { cssBundleHref } from "@remix-run/css-bundle";
import type { LinksFunction, MetaFunction } from "@remix-run/node";
import {
  Links,
  LiveReload,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  useFetchers,
  useNavigation,
} from "@remix-run/react";
import { HydrationBoundary, QueryClient, QueryClientProvider } from "@tanstack/react-query";
import nProgress from "nprogress";
import nProgressStyles from "nprogress/nprogress.css";
import { useDehydratedState } from "use-dehydrated-state";

import stylesheet from "~/tailwind.css";
import BottomNavbar from "./components/BottomNavbar";
import Header from "./components/Header";
import ThemeSwitch from "./components/ThemeSwitch";

export const meta: MetaFunction = () => [{ title: "開心團購" }];

export const links: LinksFunction = () => [
  { rel: "stylesheet", href: stylesheet },
  { rel: "stylesheet", href: nProgressStyles },
  ...(cssBundleHref ? [{ rel: "stylesheet", href: cssBundleHref }] : []),
];

export default function App() {
  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: {
            staleTime: 1000 * 60,
            // refetchOnMount: "always",
            // retry: 1,
          },
        },
      }),
  );

  const dehydratedState = useDehydratedState();

  const navigation = useNavigation();

  const fetchers = useFetchers();

  const state = useMemo<"idle" | "loading">(
    function () {
      let states = [navigation.state, ...fetchers.map((fetcher) => fetcher.state)];
      if (states.every((state) => state === "idle")) return "idle";
      return "loading";
    },
    [navigation.state, fetchers],
  );

  let progressBarTimeout = useRef<NodeJS.Timeout>();

  useEffect(() => {
    const startProgressBar = () => {
      clearTimeout(progressBarTimeout.current);
      progressBarTimeout.current = setTimeout(nProgress.start, 200);
    };

    const stopProgressBar = () => {
      clearTimeout(progressBarTimeout.current);
      nProgress.done();
    };

    if (state === "loading") {
      startProgressBar();
    } else if (state === "idle") {
      stopProgressBar();
    }
  }, [navigation.state, state]);

  useEffect(() => {
    nProgress.configure({ showSpinner: false });
  }, []);

  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width,initial-scale=1" />
        <Meta />
        <Links />
        <ColorSchemeScript defaultColorScheme="auto" />
      </head>

      <MantineProvider
        defaultColorScheme="auto"
        theme={{
          primaryColor: "lime",
        }}
      >
        <body className="mb-2 flex h-[dvh] flex-col bg-zinc-50 dark:bg-zinc-900">
          <QueryClientProvider client={queryClient}>
            <HydrationBoundary state={dehydratedState}>
              <Header>
                <ThemeSwitch />
              </Header>
              <div
                id="infiniteScrollTarget"
                className="z-0 h-full max-h-[dvh-56px] overflow-y-auto"
              >
                <Outlet />
              </div>
              <BottomNavbar />
              <ScrollRestoration />
              <Scripts />
              <LiveReload />
            </HydrationBoundary>
          </QueryClientProvider>
        </body>
      </MantineProvider>
    </html>
  );
}
