import fs from "node:fs";
import path from "node:path";
import url from "node:url";
import { createRequestHandler } from "@remix-run/express";
import type { ServerBuild } from "@remix-run/node";
import { broadcastDevReady, installGlobals } from "@remix-run/node";
import compression from "compression";
import type { RequestHandler } from "express";
import express from "express";
import morgan from "morgan";
import sourceMapSupport from "source-map-support";

sourceMapSupport.install();
installGlobals();
run();

async function run() {
  const BUILD_PATH = path.resolve("build/index.js");
  const VERSION_PATH = path.resolve("build/version.txt");
  const initialBuild = await reimportServer();

  const app = express();

  app.use((req, res, next) => {
    res.set("Strict-Transport-Security", `max-age=${60 * 60 * 24 * 365 * 100}`);

    if (req.path.endsWith("/") && req.path.length > 1) {
      const query = req.url.slice(req.path.length);
      const safepath = req.path.slice(0, -1).replace(/\/+/g, "/");
      res.redirect(301, safepath + query);
      return;
    }
    next();
  });

  app.use(compression());

  app.disable("x-powered-by");

  app.use("/build", express.static("public/build", { immutable: true, maxAge: "1y" }));

  app.use(express.static("public", { maxAge: "1h" }));

  app.use(morgan("tiny"));

  app.all(
    "*",
    process.env.NODE_ENV === "development"
      ? await createDevRequestHandler(initialBuild)
      : createRequestHandler({
          build: initialBuild,
          mode: process.env.NODE_ENV,
        }),
  );

  const port = process.env.PORT || 3000;
  app.listen(port, () => {
    console.log(`✅ app ready: http://localhost:${port}`);

    if (process.env.NODE_ENV === "development") {
      broadcastDevReady(initialBuild);
    }
  });

  async function reimportServer(): Promise<ServerBuild> {
    Object.keys(require.cache).forEach((key) => {
      if (key.startsWith(BUILD_PATH)) {
        delete require.cache[key];
      }
    });

    const stat = fs.statSync(BUILD_PATH);

    const BUILD_URL = url.pathToFileURL(BUILD_PATH).href;

    return import(BUILD_URL + "?t=" + stat.mtimeMs);
  }

  async function createDevRequestHandler(initialBuild: ServerBuild): Promise<RequestHandler> {
    let build = initialBuild;
    async function handleServerUpdate() {
      build = await reimportServer();
      broadcastDevReady(build);
    }
    const chokidar = await import("chokidar");
    chokidar
      .watch(VERSION_PATH, { ignoreInitial: true })
      .on("add", handleServerUpdate)
      .on("change", handleServerUpdate);

    return async (req, res, next) => {
      try {
        return createRequestHandler({
          build,
          mode: "development",
        })(req, res, next);
      } catch (error) {
        next(error);
      }
    };
  }
}
