const { flatRoutes } = require("remix-flat-routes");

/** @type {import('@remix-run/dev').AppConfig} */
module.exports = {
  cacheDirectory: "./node_modules/.cache/remix",
  ignoredRouteFiles: ["**/.*", "**/*.test.{ts,tsx}"],
  serverModuleFormat: "cjs",
  serverDependenciesToBundle: [/^yet-another-react-lightbox.*/],
  tailwind: true,
  postcss: true,
  routes: async (defineRoutes) => {
    return flatRoutes("routes", defineRoutes, {
      ignoredRouteFiles: [".*", "**/*.css", "**/*.test.{js,jsx,ts,tsx}", "**/__*.*"],
    });
  },
};
