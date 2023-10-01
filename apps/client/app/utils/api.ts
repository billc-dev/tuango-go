import createClient from "openapi-fetch";
import type { paths } from "types/schema";

export const serverClient = createClient<paths>({
  baseUrl: "http://127.0.0.1:5010",
});
// export const client = createClient<paths>({ baseUrl: "http://localhost:3000" });
export const client = createClient<paths>({
  baseUrl: "https://appearing-dual-investigations-mon.trycloudflare.com",
});
