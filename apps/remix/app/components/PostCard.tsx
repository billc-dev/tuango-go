import clsx from "clsx";
import { LazyLoadImage } from "react-lazy-load-image-component";
import type { components } from "types/schema";

import "react-lazy-load-image-component/src/effects/opacity.css";

import { Link } from "@remix-run/react";

interface PostCardInterface {
  post: components["schemas"]["client.paginatedPost"];
}
export const PostCard: React.FC<PostCardInterface> = ({ post }) => {
  return (
    <div className="flex max-w-[180px] transform flex-col overflow-hidden rounded-2xl bg-white antialiased shadow-sm ring-1 ring-zinc-200 transition hover:scale-[1.01] hover:shadow-lg dark:bg-zinc-800 dark:ring-zinc-700 dark:hover:shadow-gray-900">
      <Link to={`/posts?post_id=${post.id}`} className="h-[180px] w-[180px] cursor-pointer">
        <LazyLoadImage
          alt="product"
          src={post.images && post.images[0].sm}
          effect="opacity"
          className={clsx("h-[180px] w-[180px] object-cover transition-all duration-300")}
        />
      </Link>
      <div className="px-2 pb-1 pt-2">
        <Link className="cursor-pointer" to={`/posts?post_id=${post.id}`}>
          <div className="truncate">{post.title}</div>
          <div className="flex items-center justify-between">
            <div className="max-w-[50%] truncate text-xs">{post.seller?.display_name}</div>
            <div className="w-auto truncate">{`$${getProductPriceRange(post.post_items)}`}</div>
          </div>
        </Link>
        <div className="flex justify-between py-1">
          {/* <LikeButton postId={post._id} /> */}
          <Link
            aria-label="orderCount"
            className="cursor-pointer truncate"
            to={`/posts?post_id=${post.id}`}
          >
            {post.order_count ? `${post.order_count} 訂單` : ""}
          </Link>
        </div>
      </div>
    </div>
  );
};

export function getProductPriceRange(items?: components["schemas"]["ent.PostItem"][]) {
  let price: string | number = items?.[0].price ?? 0;

  if (items && items.length > 1) {
    const min = items.reduce((min, item) => Math.min(min, item.price ?? 0), items[0].price ?? 0);
    const max = items.reduce((min, item) => Math.max(min, item.price ?? 0), items[0].price ?? 0);
    if (min === max) {
      price = items[0].price ?? 0;
    } else {
      price = `${min}~$${max}`;
    }
  }

  return price;
}
