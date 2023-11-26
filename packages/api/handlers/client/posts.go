package client

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent"
	"github.com/billc-dev/tuango-go/ent/comment"
	"github.com/billc-dev/tuango-go/ent/like"
	"github.com/billc-dev/tuango-go/ent/order"
	"github.com/billc-dev/tuango-go/ent/orderitem"
	"github.com/billc-dev/tuango-go/ent/post"
	"github.com/billc-dev/tuango-go/ent/postitem"
	"github.com/billc-dev/tuango-go/ent/predicate"
	"github.com/billc-dev/tuango-go/ent/schema"
	"github.com/billc-dev/tuango-go/ent/user"
	"github.com/billc-dev/tuango-go/utils"
	"github.com/gofiber/fiber/v2"
)

// GetPosts
//
//	@Summary	Paginate posts
//	@Tags		client/posts
//	@Produce	json
//	@Param		post_num			query		number		false	"Post number"
//	@Param		status				query		post.Status	false	"Post status"
//	@Param		text				query		string		false	"Text"
//	@Param		deadline			query		string		false	"Deadline"
//	@Param		delivery_date		query		string		false	"Delivery date"
//	@Param		seller_id			query		string		false	"Seller ID"
//	@Param		page				query		number		false	"Page (0-based)"	default(0)
//	@Param		include_post_body	query		bool		false	"Include post body"
//	@Success	200					{object}	utils.InfinitePaginatedResult[paginatedPost]
//	@Failure	500					{object}	utils.HTTPError
//	@Router		/api/client/v1/posts [get]
func GetPosts(c *fiber.Ctx) error {
	m := c.Queries()

	postWhere := []predicate.Post{}
	u64, err := strconv.ParseUint(m["post_num"], 10, 64)
	if err == nil {
		postNum := int(u64)
		postWhere = append(postWhere, post.PostNumEQ(postNum))
	}

	status := post.Status(m["status"])
	err = post.StatusValidator(status)
	if status == post.StatusCompleted || status == post.StatusCanceled {
		return fiber.NewError(http.StatusInternalServerError, fmt.Sprintf(`Post status "%v" is not valid`, status))
	}
	if err != nil {
		postWhere = append(postWhere, post.StatusEQ(post.StatusOpen))
	} else {
		postWhere = append(postWhere, post.StatusEQ(status))
	}

	text := m["text"]
	if text != "" {
		postWhere = append(postWhere, post.Or(
			post.TitleContains(text),
			post.BodyContains(text),
		))
	}

	deadline := m["deadline"]
	if deadline != "" {
		postWhere = append(postWhere, post.DeadlineEQ(deadline))
	}

	deliveryDate := m["delivery_date"]
	if deliveryDate != "" {
		postWhere = append(postWhere, post.DeliveryDateEQ(deliveryDate))
	}

	sellerId := m["seller_id"]
	if sellerId != "" {
		postWhere = append(postWhere, post.SellerID(sellerId))
	}

	limit := 20
	page := 0

	if pageInt, err := strconv.Atoi(m["page"]); err == nil {
		page = int(pageInt)
	}

	offset := page * limit

	postSelect := []string{
		post.FieldPostNum, post.FieldTitle,
		post.FieldDeadline, post.FieldDeliveryDate,
		post.FieldLikeCount, post.FieldCommentCount, post.FieldOrderCount,
		post.FieldImages, post.FieldStorageType, post.FieldStatus, post.FieldCreatedAt,
	}

	if m["include_post_body"] == "true" {
		postSelect = append(postSelect, post.FieldBody)
	}

	dbquerystart := time.Now()

	posts, err := database.DBConn.Post.Query().
		Select(postSelect...).
		WithSeller(func(uq *ent.UserQuery) {
			uq.Select(user.FieldDisplayName, user.FieldPictureURL)
		}).
		WithPostItems().
		Where(postWhere...).
		Order(post.ByPostNum(sql.OrderDesc())).
		Limit(limit).
		Offset(offset).
		All(c.Context())

	if err != nil {
		return utils.Error(err, http.StatusInternalServerError, "Could not query posts")
	}
	fmt.Println("DB query", time.Since(dbquerystart))

	return c.JSON(utils.InfinitePaginatedResult[*ent.Post]{
		HasMore: len(posts) == limit,
		Data:    posts,
	})
}

type paginatedPost struct {
	ID           string           `json:"id"`
	SellerID     string           `json:"seller_id"`
	PostNum      int              `json:"post_num"`
	Title        string           `json:"title"`
	Body         string           `json:"body"`
	Deadline     string           `json:"deadline"`
	DeliveryDate string           `json:"delivery_date"`
	LikeCount    int              `json:"like_count"`
	CommentCount int              `json:"comment_count"`
	OrderCount   int              `json:"order_count"`
	Images       []schema.Image   `json:"images"`
	StorageType  post.StorageType `json:"storage_type"`
	Status       post.Status      `json:"status"`
	CreatedAt    time.Time        `json:"created_at"`
	PostItems    []ent.PostItem   `json:"post_items"`
	Seller       struct {
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
		PictureURL  string `json:"picture_url"`
	} `json:"seller"`
}

// GetPost
//
//	@Summary	Get post
//	@Tags		client/posts
//	@Produce	json
//	@Param		id	path		string	true	"Post ID"
//	@Success	200	{object}	utils.Result[normalPost]
//	@Failure	500	{object}	utils.HTTPError
//	@Router		/api/client/v1/posts/{id} [get]
func GetPost(c *fiber.Ctx) error {
	postID := c.Params("id")

	p, err := database.DBConn.Post.
		Query().
		Select(
			post.FieldPostNum, post.FieldTitle, post.FieldBody,
			post.FieldDeadline, post.FieldDeliveryDate,
			post.FieldLikeCount, post.FieldCommentCount, post.FieldOrderCount,
			post.FieldImages, post.FieldStorageType, post.FieldStatus, post.FieldCreatedAt,
		).
		Where(post.ID(postID), post.StatusIn(post.StatusOpen, post.StatusClosed)).
		WithSeller(func(uq *ent.UserQuery) {
			uq.Select(user.FieldDisplayName, user.FieldPictureURL)
		}).
		WithPostItems(func(piq *ent.PostItemQuery) {
			piq.Order(postitem.ByIdentifier())
		}).
		First(c.Context())

	if err != nil {
		return utils.Error(err, http.StatusBadRequest, "Post not found")
	}

	return c.JSON(utils.Result[*ent.Post]{
		Data: p,
	})
}

type normalPost struct {
	ID           string           `json:"id"`
	SellerID     string           `json:"seller_id"`
	PostNum      int              `json:"post_num"`
	Title        string           `json:"title"`
	Body         string           `json:"body"`
	Deadline     string           `json:"deadline"`
	DeliveryDate string           `json:"delivery_date"`
	LikeCount    int              `json:"like_count"`
	CommentCount int              `json:"comment_count"`
	OrderCount   int              `json:"order_count"`
	Images       []schema.Image   `json:"images"`
	StorageType  post.StorageType `json:"storage_type"`
	Status       post.Status      `json:"status"`
	CreatedAt    time.Time        `json:"created_at"`
	PostItems    []ent.PostItem   `json:"post_items"`
	Seller       struct {
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
		PictureURL  string `json:"picture_url"`
	} `json:"seller"`
}

func GetPostComments(c *fiber.Ctx) error {
	postID := c.Params("id")

	p, err := database.DBConn.Debug().Comment.
		Query().
		Where(comment.PostID(postID), comment.HasPostWith(post.StatusIn(post.StatusOpen, post.StatusClosed))).
		Order(comment.ByCreatedAt(sql.OrderDesc())).
		All(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Post comment not found")
	}

	return c.JSON(fiber.Map{
		"data": p,
	})
}

// GetPostOrders
//
//	@Summary	Get post orders
//	@Tags		client/posts
//	@Produce	json
//	@Security	BearerToken
//	@Param		id	path		string	true	"Post ID"
//	@Success	200	{object}	utils.Result[[]postOrder]
//	@Failure	500	{object}	utils.HTTPError
//	@Router		/api/client/v1/posts/{id}/orders [get]
func GetPostOrders(c *fiber.Ctx) error {
	postID := c.Params("id")

	o, err := database.DBConn.Order.
		Query().
		Where(
			order.PostID(postID),
			order.OrderNumNEQ(0),
			order.StatusNEQ(order.StatusCanceled),
			order.HasPostWith(post.StatusIn(post.StatusOpen, post.StatusClosed)),
		).
		WithOrderItems().
		WithUser(func(uq *ent.UserQuery) {
			uq.Select(user.FieldDisplayName, user.FieldPictureURL)
		}).
		Order(order.ByOrderNum(sql.OrderAsc())).
		All(c.Context())

	if err != nil {
		return utils.Error(err, http.StatusInternalServerError, "Could not query post orders")
	}

	return c.JSON(utils.Result[[]*ent.Order]{
		Data: o,
	})
}

type postOrder struct {
	ID            string       `json:"id"`
	UserID        string       `json:"user_id"`
	PostID        string       `json:"post_id"`
	OrderNum      int          `json:"order_num"`
	Comment       string       `json:"comment"`
	SellerComment string       `json:"seller_comment"`
	HasName       bool         `json:"has_name"`
	IsExtra       bool         `json:"is_extra"`
	Fb            bool         `json:"fb"`
	IsInStock     bool         `json:"is_in_stock"`
	Status        order.Status `json:"status"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	OrderItems    []struct {
		ID         string           `json:"id"`
		OrderID    string           `json:"order_id"`
		PostItemID string           `json:"post_item_id"`
		Identifier string           `json:"identifier"`
		Name       string           `json:"name"`
		Price      float64          `json:"price"`
		Qty        float64          `json:"qty"`
		Location   string           `json:"location"`
		HasName    bool             `json:"has_name"`
		Status     orderitem.Status `json:"status"`
	} `json:"order_items"`
	User struct {
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
		PictureURL  string `json:"picture_url"`
	} `json:"user"`
}

// LikePost
//
//	@Summary	Like post
//	@Tags		client/posts
//	@Produce	json
//	@Security	BearerToken
//	@Param		id	path		string	true	"Post ID"
//	@Success	200	{object}	utils.Result[bool]
//	@Failure	500	{object}	utils.HTTPError
//	@Router		/api/client/v1/posts/{id}/like [post]
func LikePost(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*ent.User)

	if !ok {
		return utils.Error(nil, http.StatusNotFound, "User not found")
	}

	postID := c.Params("id")

	_, err := database.DBConn.Post.
		Query().
		Where(post.ID(postID)).
		FirstID(c.Context())

	if err != nil {
		return utils.Error(err, http.StatusNotFound, "Post does not exist")
	}

	_, err = database.DBConn.Like.
		Query().
		Where(like.PostID(postID)).
		FirstID(c.Context())

	if err == nil {
		return utils.Error(err, http.StatusBadRequest, "Like already exists or could not query like")
	}

	tx, err := database.DBConn.Tx(c.Context())
	if err != nil {
		return utils.Error(err, http.StatusInternalServerError, "Could not start transaction")
	}

	err = tx.Like.
		Create().
		SetPostID(postID).
		SetUserID(u.ID).
		Exec(c.Context())
	if err != nil {
		tx.Rollback()
		return utils.Error(err, http.StatusInternalServerError, "Could not create like")
	}

	err = tx.Post.UpdateOneID(postID).AddLikeCount(1).Exec(c.Context())
	if err != nil {
		tx.Rollback()
		return utils.Error(err, http.StatusInternalServerError, "Could not increment like count")
	}

	err = tx.Commit()
	if err != nil {
		return utils.Error(err, http.StatusInternalServerError, "Could not commit transaction")
	}

	return c.JSON(utils.Result[bool]{
		Data: true,
	})
}

// UnlikePost
//
//	@Summary	Unlike post
//	@Tags		client/posts
//	@Produce	json
//	@Security	BearerToken
//	@Param		id	path		string	true	"Post ID"
//	@Success	200	{object}	utils.Result[bool]
//	@Failure	500	{object}	utils.HTTPError
//	@Router		/api/client/v1/posts/{id}/like [delete]
func UnlikePost(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*ent.User)

	if !ok {
		return fiber.NewError(http.StatusNotFound, "User not found")
	}

	postID := c.Params("id")

	_, err := database.DBConn.Post.
		Query().
		Where(post.ID(postID)).
		FirstID(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusNotFound, "Post does not exist")
	}

	tx, err := database.DBConn.Debug().Tx(c.Context())
	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not start transaction")
	}

	count, err := tx.Like.
		Delete().
		Where(
			like.UserID(u.ID),
			like.PostID(postID),
		).
		Exec(c.Context())
	if err != nil {
		log.Print(err)
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "Could not delete like")
	}
	if count != 1 {
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "Could not delete like due to count")
	}

	err = tx.Post.UpdateOneID(postID).AddLikeCount(-1).Exec(c.Context())
	if err != nil {
		log.Print(err)
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "Could not decrement post line count")
	}

	err = tx.Commit()
	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not commit order")
	}

	return c.JSON(fiber.Map{})
}
