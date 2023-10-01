import type { components } from "types/schema";

export const getStorageTypeLabel = (storageType?: components["schemas"]["post.StorageType"]) => {
  switch (storageType) {
    case "roomTemp":
      return "常溫";
    case "farmGoods":
      return "農產品";
    case "refrigerated":
      return "冷藏";
    case "frozen":
      return "冷凍";
    default:
      return "";
  }
};
