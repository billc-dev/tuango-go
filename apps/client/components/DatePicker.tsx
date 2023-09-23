"use client";

import "dayjs/locale/zh-tw";

import { DatePickerInput } from "@mantine/dates";

type DatePickerProps = React.ComponentProps<typeof DatePickerInput>;

export const DatePicker = (props: DatePickerProps) => {
  return (
    <DatePickerInput
      {...props}
      valueFormat="YYYY-MM-DD dddd"
      allowDeselect
      clearable
    />
  );
};
