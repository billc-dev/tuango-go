import { useQuery } from "@tanstack/react-query";

import { useEnv } from "~/root";
import { client } from "~/utils/api";

export const useUser = () => {
  const { authenticated } = useEnv();

  return useQuery({
    enabled: authenticated,
    queryKey: ["user"],
    queryFn: async () => {
      const { data, error } = await client.GET("/api/client/v1/user", {});
      if (error) {
        throw new Error(error);
      }
      return { ...data };
    },
  });
};
