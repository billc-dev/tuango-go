import "@mantine/carousel/styles.css";
import "@mantine/core/styles.css";
import "@mantine/dates/styles.css";

import { createContext, useContext, useEffect, useMemo, useRef, useState } from "react";
import { ColorSchemeScript, MantineProvider } from "@mantine/core";
import { cssBundleHref } from "@remix-run/css-bundle";
import type { DataFunctionArgs, LinksFunction, MetaFunction } from "@remix-run/node";
import { json } from "@remix-run/node";
import {
  Links,
  LiveReload,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  useFetchers,
  useLoaderData,
  useNavigation,
} from "@remix-run/react";
import {
  dehydrate,
  HydrationBoundary,
  QueryClient,
  QueryClientProvider,
} from "@tanstack/react-query";
import nProgress from "nprogress";
import nProgressStyles from "nprogress/nprogress.css";
import { useDehydratedState } from "use-dehydrated-state";

import stylesheet from "~/tailwind.css";
import BottomNavbar from "./components/BottomNavbar";
import Header from "./components/Header";
import ReactToaster from "./components/ReactToaster";
import ThemeSwitch from "./components/ThemeSwitch";
import { serverClient } from "./utils/api";
import { getSession } from "./utils/session.server";

export const meta: MetaFunction = () => [{ title: "開心團購" }];

export const links: LinksFunction = () => [
  { rel: "stylesheet", href: stylesheet },
  { rel: "stylesheet", href: nProgressStyles },
  ...(cssBundleHref ? [{ rel: "stylesheet", href: cssBundleHref }] : []),
];

export const shouldRevalidate = () => false;

export const loader = async ({ request }: DataFunctionArgs) => {
  const session = await getSession(request.headers.get("Cookie"));

  const token = session.get("token");
  const authenticated = Boolean(token);
  const post_id = new URL(request.url).searchParams.get("post_id");

  const queryClient = new QueryClient();

  if (authenticated && !post_id) {
    await queryClient.prefetchQuery({
      // eslint-disable-next-line @tanstack/query/exhaustive-deps
      queryKey: ["user"],
      queryFn: async () => {
        const { data, error } = await serverClient.GET("/api/client/v1/user", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        if (error) {
          throw new Error(error);
        }
        return { ...data };
      },
    });
  }

  return json({
    LINE_CALLBACK_URL: process.env.LINE_CALLBACK_URL,
    dehydratedState: dehydrate(queryClient),
    authenticated,
  });
};

const EnvContext = createContext({ authenticated: false, LINE_CALLBACK_URL: "" });
export const useEnv = () => useContext(EnvContext);

export default function App() {
  const { LINE_CALLBACK_URL, authenticated } = useLoaderData<typeof loader>();

  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: {
            staleTime: 1000 * 60,
            refetchOnMount: "always",
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

      <EnvContext.Provider value={{ authenticated, LINE_CALLBACK_URL }}>
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
                <ReactToaster />
                <div id="infiniteScrollTarget" className="h-full max-h-[dvh-56px] overflow-y-auto">
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
      </EnvContext.Provider>
    </html>
  );
}
