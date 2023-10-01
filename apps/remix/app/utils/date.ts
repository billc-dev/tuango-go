import dayjs from "dayjs";

import "dayjs/locale/zh-tw";

import localizedFormat from "dayjs/plugin/localizedFormat";
import relativeTime from "dayjs/plugin/relativeTime";

dayjs.extend(localizedFormat);
dayjs.extend(relativeTime);
dayjs.locale("zh-tw");

export const getFullDateFromNow = (date?: string) => {
  return `${dayjs(date).format("LLLL")} ${dayjs(date).fromNow()}`;
};

export const date = (date: string) => {
  return dayjs(date).format("LLLL").slice(0, -6);
};
