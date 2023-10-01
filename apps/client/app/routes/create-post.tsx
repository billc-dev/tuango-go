import { Fragment } from "react";
import { PhotoIcon, PlusIcon, TrashIcon } from "@heroicons/react/24/outline";
import {
  Button,
  Checkbox,
  FileButton,
  Grid,
  Group,
  List,
  NumberInput,
  Select,
  Text,
  Textarea,
  TextInput,
} from "@mantine/core";
import { DatesProvider } from "@mantine/dates";
import { useForm } from "@mantine/form";
import { randomId } from "@mantine/hooks";
import { type MetaFunction } from "@remix-run/node";
import dayjs from "dayjs";

import { DatePicker } from "~/components/DatePicker";

export const meta: MetaFunction = () => [{ title: "新增訂單 - 開心團購" }];

export default function Route() {
  const form = useForm({
    initialValues: {
      title: "",
      storage_type: "roomTemp",
      deadline: null,
      delivery_date: null,
      body: "",
      items: [
        {
          key: randomId(),
          name: "",
          price: "",
          qty: 100,
        },
      ],
      confirmed: false,
    },
  });

  return (
    <div className="mx-2 my-2 space-y-2">
      <TextInput
        label="團購主題"
        placeholder="團購主題"
        size="md"
        required
        description="👉 流水編號和團主名稱不用寫"
        inputWrapperOrder={["label", "input", "description", "error"]}
        {...form.getInputProps("title")}
      />
      <Select
        label="儲存方式"
        required
        size="md"
        data={[
          { label: "常溫", value: "roomTemp" },
          { label: "農產品", value: "farmGoods" },
          { label: "冷藏 ❄️", value: "refrigerated" },
          { label: "冷凍 🧊", value: "frozen" },
        ]}
        {...form.getInputProps("storage_type")}
      />
      <DatesProvider settings={{ locale: "zh-tw" }}>
        <DatePicker
          label="結單日"
          size="md"
          placeholder="結單日"
          autoFocus={false}
          {...form.getInputProps("deadline")}
        />
        <DatePicker
          label="到貨日"
          size="md"
          placeholder="到貨日"
          excludeDate={(date) => {
            return dayjs(date).day() === 0;
          }}
          {...form.getInputProps("delivery_date")}
        />
      </DatesProvider>
      <Textarea
        label="團購內容"
        placeholder="團購內容"
        size="md"
        autosize
        minRows={5}
        maxRows={10}
        required
        {...form.getInputProps("body")}
      />
      {form.values.items.map((item, index) => (
        <Fragment key={item.key}>
          <TextInput
            autoFocus={index > 0}
            label={`${String.fromCharCode("A".charCodeAt(0) + index)}. 商品名稱`}
            placeholder="商品名稱"
            size="md"
            required
            {...form.getInputProps(`items.${index}.name`)}
          />
          <Grid>
            <Grid.Col span={6}>
              <NumberInput
                size="md"
                prefix="$"
                label="價格"
                placeholder="價格"
                required
                min={0}
                hideControls
                {...form.getInputProps(`items.${index}.price`)}
              />
            </Grid.Col>
            <Grid.Col span={6}>
              <NumberInput
                size="md"
                label="數量"
                placeholder="數量"
                required
                min={0}
                hideControls
                {...form.getInputProps(`items.${index}.qty`)}
              />
            </Grid.Col>
          </Grid>
          <Group justify="space-between" mt="sm">
            <Button
              leftSection={<PlusIcon className="h-6 w-6" />}
              variant="filled"
              size="compact-md"
              color="gray"
              onClick={() =>
                form.insertListItem(
                  "items",
                  {
                    key: randomId(),
                    name: "",
                    price: "",
                    qty: 100,
                  },
                  index + 1,
                )
              }
            >
              插入商品
            </Button>
            <Button
              leftSection={<TrashIcon className="h-5 w-5" />}
              variant="filled"
              size="compact-md"
              color="red"
              onClick={() => form.removeListItem("items", index)}
            >
              刪除
            </Button>
          </Group>
        </Fragment>
      ))}
      <Button
        leftSection={<PlusIcon className="h-7 w-7" />}
        variant="filled"
        color="gray"
        my="md"
        fullWidth
        onClick={() =>
          form.insertListItem("items", {
            key: randomId(),
            name: "",
            price: "",
            qty: 100,
          })
        }
      >
        新增商品
      </Button>
      <FileButton onChange={() => {}} accept="image/png,image/jpeg" multiple>
        {(props) => (
          <Button
            leftSection={<PhotoIcon className="h-6 w-6" />}
            variant="filled"
            color="gray"
            my="md"
            fullWidth
            {...props}
          >
            新增照片
          </Button>
        )}
      </FileButton>
      <Text className="inline" size="md">
        ⚠️ 開單注意事項:
      </Text>
      <List>
        <List.Item>結單前，團員有權利取消訂單</List.Item>
        <List.Item>平台費為6%，待認購為10%</List.Item>
      </List>
      <Checkbox size="md" label="我已閱讀並同意開單注意事項" {...form.getInputProps("confirmed")} />
    </div>
  );
}
