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

export const getOrderStatusLabel = (orderStatus?: components["schemas"]["order.Status"]) => {
  switch (orderStatus) {
    case "delivered":
      return "已到貨 🚚";
    case "completed":
      return "已取貨 ✅";
    case "missing":
      return "尋貨中 🔍";
    default:
      return "";
  }
};
