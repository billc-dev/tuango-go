/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export interface paths {
  "/api/admin/v1/posts": {
    /** Paginate posts */
    get: {
      parameters: {
        query?: {
          /** @description Post number */
          post_num?: number;
          /** @description Post status */
          status?: "open" | "closed" | "completed" | "canceled";
          /** @description Text */
          text?: string;
          /** @description Deadline */
          deadline?: string;
          /** @description Delivery date */
          delivery_date?: string;
          /** @description Seller ID */
          seller_id?: string;
          /** @description Page (0-based) */
          page?: number;
        };
      };
      responses: {
        /** @description OK */
        200: {
          content: {
            "application/json": string;
          };
        };
        /** @description Internal Server Error */
        500: {
          content: {
            "application/json": string;
          };
        };
      };
    };
    /** Create post */
    post: {
      /** @description Post body */
      requestBody: {
        content: {
          "application/json": components["schemas"]["admin.createPostForm"];
        };
      };
      responses: {
        /** @description OK */
        200: {
          content: {
            "application/json": string;
          };
        };
        /** @description Internal Server Error */
        500: {
          content: {
            "application/json": string;
          };
        };
      };
    };
  };
  "/api/client/v1/orders": {
    /** Create order */
    post: {
      /** @description Order form */
      requestBody: {
        content: {
          "application/json": components["schemas"]["client.orderForm"];
        };
      };
      responses: {
        /** @description OK */
        200: {
          content: {
            "application/json": components["schemas"]["utils.Result-client_createOrderData"];
          };
        };
        /** @description Internal Server Error */
        500: {
          content: {
            "application/json": string;
          };
        };
      };
    };
  };
  "/api/client/v1/orders/{id}": {
    /** Cancel order */
    delete: {
      parameters: {
        path: {
          /** @description Order ID */
          id: string;
        };
      };
      responses: {
        /** @description OK */
        200: {
          content: {
            "application/json": components["schemas"]["utils.Result-bool"];
          };
        };
        /** @description Internal Server Error */
        500: {
          content: {
            "application/json": string;
          };
        };
      };
    };
  };
  "/api/client/v1/posts": {
    /** Paginate posts */
    get: {
      parameters: {
        query?: {
          /** @description Post number */
          post_num?: number;
          /** @description Text */
          text?: string;
          /** @description Deadline */
          deadline?: string;
          /** @description Delivery date */
          delivery_date?: string;
          /** @description Seller ID */
          seller_id?: string;
          /** @description Page (0-based) */
          page?: number;
          /** @description Include post body */
          include_post_body?: boolean;
        };
      };
      responses: {
        /** @description OK */
        200: {
          content: {
            "application/json": components["schemas"]["utils.InfinitePaginatedResult-client_paginatedPost"];
          };
        };
        /** @description Internal Server Error */
        500: {
          content: {
            "application/json": string;
          };
        };
      };
    };
  };
  "/api/client/v1/posts/{id}": {
    /** Get post */
    get: {
      parameters: {
        path: {
          /** @description Post ID */
          id: string;
        };
      };
      responses: {
        /** @description OK */
        200: {
          content: {
            "application/json": components["schemas"]["utils.Result-client_normalPost"];
          };
        };
        /** @description Internal Server Error */
        500: {
          content: {
            "application/json": string;
          };
        };
      };
    };
  };
  "/api/client/v1/posts/{id}/like": {
    /** Like post */
    post: {
      parameters: {
        path: {
          /** @description Post ID */
          id: string;
        };
      };
      responses: {
        /** @description OK */
        200: {
          content: {
            "application/json": components["schemas"]["utils.Result-bool"];
          };
        };
        /** @description Internal Server Error */
        500: {
          content: {
            "application/json": string;
          };
        };
      };
    };
    /** Unlike post */
    delete: {
      parameters: {
        path: {
          /** @description Post ID */
          id: string;
        };
      };
      responses: {
        /** @description OK */
        200: {
          content: {
            "application/json": components["schemas"]["utils.Result-bool"];
          };
        };
        /** @description Internal Server Error */
        500: {
          content: {
            "application/json": string;
          };
        };
      };
    };
  };
  "/api/client/v1/posts/{id}/orders": {
    /** Get post orders */
    get: {
      parameters: {
        path: {
          /** @description Post ID */
          id: string;
        };
      };
      responses: {
        /** @description OK */
        200: {
          content: {
            "application/json": components["schemas"]["utils.Result-array_client_postOrder"];
          };
        };
        /** @description Internal Server Error */
        500: {
          content: {
            "application/json": string;
          };
        };
      };
    };
  };
  "/api/client/v1/user": {
    /** Get user */
    get: {
      responses: {
        /** @description OK */
        200: {
          content: {
            "application/json": components["schemas"]["utils.Result-client_getUserData"];
          };
        };
        /** @description Bad Request */
        400: {
          content: {
            "application/json": string;
          };
        };
        /** @description Unauthorized */
        401: {
          content: {
            "application/json": string;
          };
        };
        /** @description Internal Server Error */
        500: {
          content: {
            "application/json": string;
          };
        };
      };
    };
  };
  "/api/client/v1/user/likes": {
    /** Get likes */
    get: {
      responses: {
        /** @description OK */
        200: {
          content: {
            "application/json": components["schemas"]["utils.Result-array_ent_Like"];
          };
        };
        /** @description Unauthorized */
        401: {
          content: {
            "application/json": string;
          };
        };
        /** @description Internal Server Error */
        500: {
          content: {
            "application/json": string;
          };
        };
      };
    };
  };
  "/api/client/v1/user/login/line/{code}": {
    /** Line login */
    post: {
      parameters: {
        query: {
          /** @description Line login redirect uri */
          redirect_uri: string;
        };
        path: {
          /** @description Line login code */
          code: string;
        };
      };
      responses: {
        /** @description OK */
        200: {
          content: {
            "application/json": components["schemas"]["utils.Result-string"];
          };
        };
        /** @description Internal Server Error */
        500: {
          content: {
            "application/json": string;
          };
        };
      };
    };
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    "admin.createPostForm": {
      body?: string;
      deadline?: string;
      delivery_date?: string;
      images?: components["schemas"]["schema.Image"][];
      is_in_stock?: boolean;
      items?: {
          name?: string;
          price?: number;
          stock?: number;
        }[];
      post_num?: number;
      seller_id?: string;
      storage_type?: components["schemas"]["post.StorageType"];
      title?: string;
    };
    "client.createOrderData": {
      /** @description Comment holds the value of the "comment" field. */
      comment?: string;
      /** @description CreatedAt holds the value of the "created_at" field. */
      created_at?: string;
      /** @description Fb holds the value of the "fb" field. */
      fb?: boolean;
      /** @description HasName holds the value of the "has_name" field. */
      has_name?: boolean;
      /** @description ID of the ent. */
      id?: string;
      /** @description IsExtra holds the value of the "is_extra" field. */
      is_extra?: boolean;
      /** @description IsInStock holds the value of the "is_in_stock" field. */
      is_in_stock?: boolean;
      /** @description OrderNum holds the value of the "order_num" field. */
      order_num?: number;
      post?: {
        body?: string;
        comment_count?: number;
        created_at?: string;
        deadline?: string;
        delivery_date?: string;
        id?: string;
        images?: components["schemas"]["schema.Image"][];
        like_count?: number;
        order_count?: number;
        post_items?: components["schemas"]["ent.PostItem"][];
        post_num?: number;
        seller?: {
          display_name?: string;
          id?: string;
          picture_url?: string;
        };
        seller_id?: string;
        status?: components["schemas"]["post.Status"];
        storage_type?: string;
        title?: string;
      };
      /** @description PostID holds the value of the "post_id" field. */
      post_id?: string;
      /** @description SellerComment holds the value of the "seller_comment" field. */
      seller_comment?: string;
      status?: components["schemas"]["order.Status"];
      /** @description UpdatedAt holds the value of the "updated_at" field. */
      updated_at?: string;
      /** @description UserID holds the value of the "user_id" field. */
      user_id?: string;
    };
    "client.getUserData": {
      display_name?: string;
      id?: string;
      notified?: boolean;
      pickup_num?: number;
      picture_url?: string;
      role?: components["schemas"]["user.Role"];
      status?: components["schemas"]["user.Status"];
    };
    "client.normalPost": {
      body?: string;
      comment_count?: number;
      created_at?: string;
      deadline?: string;
      delivery_date?: string;
      id?: string;
      images?: components["schemas"]["schema.Image"][];
      like_count?: number;
      order_count?: number;
      post_items?: components["schemas"]["ent.PostItem"][];
      post_num?: number;
      seller?: {
        display_name?: string;
        id?: string;
        picture_url?: string;
      };
      seller_id?: string;
      status?: string;
      storage_type?: string;
      title?: string;
    };
    "client.orderForm": {
      comment?: string;
      order?: {
        [key: string]: number;
      };
      postId?: string;
      sum?: number;
    };
    "client.paginatedPost": {
      body?: string;
      comment_count?: number;
      created_at?: string;
      deadline?: string;
      delivery_date?: string;
      id?: string;
      images?: components["schemas"]["schema.Image"][];
      like_count?: number;
      order_count?: number;
      post_items?: components["schemas"]["ent.PostItem"][];
      post_num?: number;
      seller?: {
        display_name?: string;
        id?: string;
        picture_url?: string;
      };
      seller_id?: string;
      status?: string;
      storage_type?: string;
      title?: string;
    };
    "client.postOrder": {
      comment?: string;
      created_at?: string;
      fb?: boolean;
      has_name?: boolean;
      id?: string;
      is_extra?: boolean;
      is_in_stock?: boolean;
      order_items?: {
          has_name?: boolean;
          id?: string;
          identifier?: string;
          location?: string;
          name?: string;
          order_id?: string;
          post_item_id?: string;
          price?: number;
          qty?: number;
          status?: components["schemas"]["orderitem.Status"];
        }[];
      order_num?: number;
      post_id?: string;
      seller_comment?: string;
      /** @description Status holds the value of the "status" field. */
      status?: string;
      updated_at?: string;
      user?: {
        display_name?: string;
        id?: string;
        picture_url?: string;
      };
      user_id?: string;
    };
    "ent.Like": {
      /** @description CreatedAt holds the value of the "created_at" field. */
      created_at?: string;
      /** @description ID of the ent. */
      id?: string;
      /** @description PostID holds the value of the "post_id" field. */
      post_id?: string;
      /** @description UserID holds the value of the "user_id" field. */
      user_id?: string;
    };
    "ent.PostItem": {
      /** @description ID of the ent. */
      id?: string;
      /** @description Identifier holds the value of the "identifier" field. */
      identifier?: string;
      /** @description Name holds the value of the "name" field. */
      name?: string;
      /** @description PostID holds the value of the "post_id" field. */
      post_id?: string;
      /** @description Price holds the value of the "price" field. */
      price?: number;
      /** @description Stock holds the value of the "stock" field. */
      stock?: number;
    };
    /** @description Status holds the value of the "status" field. */
    "order.Status": string;
    "orderitem.Status": string;
    "post.Status": string;
    "post.StorageType": string;
    "schema.Image": {
      lg?: string;
      md?: string;
      sm?: string;
    };
    "user.Role": string;
    "user.Status": string;
    "utils.InfinitePaginatedResult-client_paginatedPost": {
      data?: components["schemas"]["client.paginatedPost"][];
      has_more?: boolean;
    };
    "utils.Result-array_client_postOrder": {
      data?: components["schemas"]["client.postOrder"][];
    };
    "utils.Result-array_ent_Like": {
      data?: components["schemas"]["ent.Like"][];
    };
    "utils.Result-bool": {
      data?: boolean;
    };
    "utils.Result-client_createOrderData": {
      data?: components["schemas"]["client.createOrderData"];
    };
    "utils.Result-client_getUserData": {
      data?: components["schemas"]["client.getUserData"];
    };
    "utils.Result-client_normalPost": {
      data?: components["schemas"]["client.normalPost"];
    };
    "utils.Result-string": {
      data?: string;
    };
  };
  responses: never;
  parameters: never;
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export type operations = Record<string, never>;
