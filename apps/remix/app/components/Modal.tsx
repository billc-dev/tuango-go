import { useEffect, useState } from "react";
import { ArrowLeftIcon, ChevronLeftIcon, ChevronRightIcon } from "@heroicons/react/24/outline";
import { Carousel } from "@mantine/carousel";
import { ActionIcon, Avatar, Modal as MantineModal } from "@mantine/core";
import { useDisclosure } from "@mantine/hooks";
import { useSearchParams } from "@remix-run/react";
import { useQuery } from "@tanstack/react-query";
// import Lightbox from "react-image-lightbox";
import { LazyLoadImage } from "react-lazy-load-image-component";
import Lightbox from "yet-another-react-lightbox";
import Counter from "yet-another-react-lightbox/plugins/counter";

import "yet-another-react-lightbox/styles.css";
import "yet-another-react-lightbox/plugins/counter.css";

import Thumbnails from "yet-another-react-lightbox/plugins/thumbnails";
import Zoom from "yet-another-react-lightbox/plugins/zoom";

// import "react-image-lightbox/style.css";
import "yet-another-react-lightbox/plugins/thumbnails.css";

import { client } from "~/utils/api";
import { date, getFullDateFromNow } from "~/utils/date";
import { getStorageTypeLabel } from "~/utils/text";

import "react-lazy-load-image-component/src/effects/opacity.css";

export const Modal = () => {
  const [searchParams, setSearchParams] = useSearchParams();

  const postId = searchParams.get("post_id");

  const [opened, { close, open }] = useDisclosure(Boolean(postId), {
    onClose: () => {
      searchParams.delete("post_id");
      setSearchParams(searchParams);
    },
  });

  const query = useQuery({
    enabled: Boolean(postId),
    queryKey: ["post", postId],
    queryFn: async () => {
      const { data, error } = await client.GET("/api/client/v1/posts/{id}", {
        params: {
          path: {
            id: postId ?? "",
          },
        },
      });
      if (error) {
        throw new Error(error.message);
      }
      return { ...data };
    },
  });

  useEffect(() => {
    if (query.data?.data) {
      open();
    } else {
      close();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [query.data?.data]);

  const post = query.data?.data;

  const [imageIndex, setImageIndex] = useState(0);
  const [openLightbox, setOpenLightbox] = useState(false);

  return (
    <MantineModal.Root
      opened={opened}
      onClose={close}
      fullScreen
      radius={0}
      // transitionProps={{
      //   transition: "fade",
      // }}
    >
      <MantineModal.Overlay />
      <MantineModal.Content>
        <MantineModal.Header className="flex items-center !justify-start border-b border-zinc-200 !px-2 !py-3.5 shadow-sm dark:border-zinc-700 dark:bg-zinc-800">
          <ActionIcon onClick={close} variant="subtle" color="gray" size="md">
            <ArrowLeftIcon className="h-6 w-6 text-black dark:text-white" />
          </ActionIcon>
          <MantineModal.Title className="line-clamp-1 !pl-2 !text-xl">
            #{query.data?.data?.post_num} {query.data?.data?.title}
          </MantineModal.Title>
        </MantineModal.Header>
        <MantineModal.Body className="mt-2">
          <div className="flex">
            <Avatar
              src={query.data?.data?.seller?.picture_url}
              alt={query.data?.data?.seller?.display_name}
            />
            <div className="ml-1 flex flex-col pl-2">
              <div className="line-clamp-1 text-sm">{query.data?.data?.seller?.display_name}</div>
              <div className="line-clamp-1 text-xs text-zinc-400">
                {getFullDateFromNow(query.data?.data?.created_at)}
              </div>
            </div>
          </div>
          <Carousel
            className="relative -mx-4 mt-2"
            onSlideChange={(index) => setImageIndex(index)}
            // loop
            withIndicators
            speed={30}
            height={300}
            controlsOffset="xs"
            previousControlIcon={<ChevronLeftIcon className="h-8 w-8" />}
            nextControlIcon={<ChevronRightIcon className="h-8 w-8" />}
            classNames={{
              indicators: "absolute px-2 ",
              indicator: "!bg-zinc-600 dark:!bg-white",
              control:
                "!bg-transparent !outline-none !border-none text-black dark:text-white !shadow-none",
            }}
          >
            {post?.images?.map((image, index) => (
              <Carousel.Slide key={index} className="flex items-center justify-center">
                <LazyLoadImage
                  src={image.md}
                  visibleByDefault={imageIndex + 1 === index}
                  effect="opacity"
                  placeholder={<div className="h-72 w-full bg-zinc-300 dark:bg-zinc-600" />}
                  className={`max-h-72 object-contain transition-all duration-300`}
                  onClick={() => setOpenLightbox(true)}
                />
              </Carousel.Slide>
            ))}
          </Carousel>
          {openLightbox && (
            <Lightbox
              open={openLightbox}
              close={() => setOpenLightbox(false)}
              slides={post?.images?.map((image) => ({
                src: image.md ?? "",
              }))}
              index={imageIndex}
              plugins={[Counter, Thumbnails, Zoom]}

              // reactModalProps={{
              //   isOpen: openLightbox,
              //   className: openLightbox ? "" : "pointer-events-none",
              // }}
              // mainSrc={post?.images?.[imageIndex].md ?? ""}
              // nextSrc={post?.images?.[(imageIndex + 1) % post?.images?.length].md}
              // prevSrc={
              //   post?.images?.[(imageIndex + post?.images?.length - 1) % post?.images?.length].md
              // }
              // onCloseRequest={() => setOpenLightbox(false)}
              // onMovePrevRequest={() =>
              //   setImageIndex(
              //     (imageIndex + (post?.images?.length ?? 1) - 1) % (post?.images?.length ?? 1),
              //   )
              // }
              // onMoveNextRequest={() => setImageIndex((imageIndex + 1) % (post?.images?.length ?? 1))}
            />
          )}
          <div className={`py-4`}>
            <p className="font-bold">
              #{post?.post_num} {post?.title} #{post?.seller?.display_name}
            </p>
            <p>❤️ #結單日: {post?.deadline ? date(post?.deadline) : "成團通知"}</p>
            <p>💚 #到貨日: {post?.delivery_date ? date(post?.delivery_date) : "貨到通知"}</p>
            <p>儲存方式: {getStorageTypeLabel(post?.storage_type)}</p>
            <p className={`whitespace-pre-line pt-4`}>{post?.body?.trim()}</p>
          </div>
        </MantineModal.Body>
      </MantineModal.Content>
    </MantineModal.Root>
  );
};
