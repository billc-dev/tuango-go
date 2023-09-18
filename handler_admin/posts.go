package handler_admin

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent"
	"github.com/billc-dev/tuango-go/ent/deliver"
	"github.com/billc-dev/tuango-go/ent/order"
	"github.com/billc-dev/tuango-go/ent/post"
	"github.com/billc-dev/tuango-go/ent/postitem"
	"github.com/billc-dev/tuango-go/ent/predicate"
	"github.com/billc-dev/tuango-go/ent/schema"
	"github.com/billc-dev/tuango-go/ent/user"
	"github.com/billc-dev/tuango-go/handler_seller"
	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	m := c.Queries()

	postWhere := []predicate.Post{}

	if postNum, err := strconv.Atoi(m["postNum"]); err == nil {
		postWhere = append(postWhere, post.PostNumEQ(postNum))
	}

	status := post.Status(m["status"])
	if err := post.StatusValidator(status); err != nil {
		return fiber.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(`Post status "%v" is not valid`, status),
		)
	} else {
		postWhere = append(postWhere, post.StatusEQ(status))
	}

	if text := m["text"]; text != "" {
		postWhere = append(postWhere, post.Or(
			post.TitleContains(text),
			post.BodyContains(text),
		))
	}

	if deadline := m["deadline"]; deadline != "" {
		postWhere = append(postWhere, post.DeadlineEQ(deadline))
	}

	if deliveryDate := m["deliveryDate"]; deliveryDate != "" {
		postWhere = append(postWhere, post.DeliveryDateEQ(deliveryDate))
	}

	sellerId := m["sellerId"]
	if sellerId != "" {
		postWhere = append(postWhere, post.SellerID(sellerId))
	}

	limit := 20
	page := 0

	pageInt, err := strconv.Atoi(m["page"])
	if err == nil {
		page = int(pageInt)
	}

	offset := page * limit

	posts, err := database.DBConn.Post.Query().
		Select(
			post.FieldPostNum, post.FieldTitle,
			post.FieldDeadline, post.FieldDeliveryDate,
			post.FieldOrderCount, post.FieldStorageType,
			post.FieldStatus, post.FieldCreatedAt,
		).
		WithSeller(func(uq *ent.UserQuery) {
			uq.Select(user.FieldDisplayName, user.FieldPictureURL)
		}).
		Where(postWhere...).
		Order(post.ByPostNum(sql.OrderDesc())).
		Limit(limit).
		Offset(offset).
		All(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query posts")
	}

	count, err := database.DBConn.Post.Query().Where(postWhere...).Count(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query post count")
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"posts":       posts,
			"count":       count,
			"hasNextPage": len(posts) == limit,
		},
	})
}

func GetPostsByDate(c *fiber.Ctx) error {
	postWhere := []predicate.Post{
		post.HasPostOrdersWith(order.StatusIn(order.StatusOrdered, order.StatusConfirmed)),
	}

	if date := c.Params("date"); date == "" {
		postWhere = append(postWhere, post.DeliveryDate(time.Now().Format("2006-01-02")))
	} else {
		postWhere = append(postWhere, post.DeliveryDate(date))
	}

	posts, err := database.DBConn.Debug().Post.Query().
		Select(post.FieldPostNum, post.FieldTitle).
		WithSeller(func(uq *ent.UserQuery) {
			uq.Select(user.FieldDisplayName)
		}).
		Where(postWhere...).
		All(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query posts")
	}

	return c.JSON(fiber.Map{
		"data": posts,
	})

}

type createPostForm struct {
	SellerID     string           `json:"seller_id"`
	PostNum      int              `json:"post_num"`
	Title        string           `json:"title"`
	Body         string           `json:"body"`
	Deadline     string           `json:"deadline"`
	DeliveryDate string           `json:"delivery_date"`
	IsInStock    bool             `json:"is_in_stock"`
	StorageType  post.StorageType `json:"storage_type"`
	Images       []schema.Image   `json:"images"`
	Items        []struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Stock float64 `json:"stock"`
	} `json:"items"`
}

func CreatePost(c *fiber.Ctx) error {
	postForm := new(createPostForm)

	if err := c.BodyParser(postForm); err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusBadRequest, "Could not parse order form")
	}

	tx, err := database.DBConn.Debug().Tx(c.Context())
	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not start transaction")
	}

	if postForm.PostNum != 0 {
		post, err := database.DBConn.Post.Query().
			Where(
				post.PostNum(postForm.PostNum),
				post.StatusNEQ(post.StatusCanceled),
			).
			Select(post.FieldPostNum).
			Order(post.ByPostNum(sql.OrderDesc())).
			First(c.Context())

		if err != nil {
			log.Print(err)
			return fiber.NewError(http.StatusInternalServerError, "Could not query previous post")
		}

		if *post.PostNum != 0 {
			return fiber.NewError(http.StatusBadRequest, "PostNum is duplicated")
		}
	} else {
		previousPost, err := database.DBConn.Post.Query().
			Where(post.StatusNEQ(post.StatusCanceled)).
			Select(post.FieldPostNum).
			Order(post.ByPostNum(sql.OrderDesc())).
			First(c.Context())

		if err != nil {
			log.Print(err)
			return fiber.NewError(http.StatusInternalServerError, "Could not query previous post")
		}

		postForm.PostNum = *previousPost.PostNum + 1
	}

	newPost, err := tx.Post.Create().
		SetSellerID(postForm.SellerID).
		SetPostNum(postForm.PostNum).
		SetTitle(postForm.Title).
		SetBody(postForm.Body).
		SetStorageType(postForm.StorageType).
		SetDeadline(postForm.Deadline).
		SetDeliveryDate(postForm.DeliveryDate).
		SetImages(postForm.Images).
		SetIsInStock(postForm.IsInStock).
		Save(c.Context())

	if err != nil {
		log.Print(err)
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "Could not create post")
	}

	if len(postForm.Items) > len(handler_seller.Alphabets) {
		tx.Rollback()
		return fiber.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf("Post items length is more than %v", len(handler_seller.Alphabets)),
		)
	}

	postItems := []*ent.PostItemCreate{}

	for index, item := range postForm.Items {
		postItems = append(postItems,
			tx.PostItem.Create().
				SetPostID(newPost.ID).
				SetIdentifier(handler_seller.Alphabets[index]).
				SetName(item.Name).
				SetPrice(item.Price).
				SetStock(item.Stock),
		)
	}

	if len(postItems) == 0 {
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "No post items to create")
	}

	err = tx.PostItem.CreateBulk(postItems...).Exec(c.Context())
	if err != nil {
		log.Print(err)
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "Could not create post")
	}

	err = tx.Commit()
	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not commit transaction")
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}

type updatePostForm struct {
	Title        string           `json:"title"`
	Body         string           `json:"body"`
	Deadline     string           `json:"deadline"`
	DeliveryDate string           `json:"delivery_date"`
	StorageType  post.StorageType `json:"storage_type"`
	IsInStock    bool             `json:"is_in_stock"`
	Images       []schema.Image   `json:"images"`
	Items        []struct {
		ID    string  `json:"id"`
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Stock float64 `json:"stock"`
	} `json:"items"`
}

func UpdatePost(c *fiber.Ctx) error {
	postID := c.Params("id")

	postForm := new(updatePostForm)

	if err := c.BodyParser(postForm); err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusBadRequest, "Could not parse order form")
	}

	tx, err := database.DBConn.Debug().Tx(c.Context())
	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not start transaction")
	}

	err = tx.Post.UpdateOneID(postID).
		SetTitle(postForm.Title).
		SetBody(postForm.Body).
		SetDeadline(postForm.Deadline).
		SetDeliveryDate(postForm.DeliveryDate).
		SetStorageType(postForm.StorageType).
		SetIsInStock(postForm.IsInStock).
		SetImages(postForm.Images).
		Exec(c.Context())

	if err != nil {
		log.Print(err)
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "Could not update post")
	}

	postItems := []*ent.PostItemCreate{}

	for index, item := range postForm.Items {
		log.Print(item)
		if item.ID != "" {
			err := tx.PostItem.UpdateOneID(item.ID).
				SetName(item.Name).
				SetPrice(item.Price).
				SetStock(item.Stock).
				Exec(c.Context())
			if err != nil {
				log.Print(err)
				tx.Rollback()
				return fiber.NewError(http.StatusInternalServerError, "Could not update post item")
			}
			continue
		}
		postItems = append(postItems,
			tx.PostItem.Create().
				SetPostID(postID).
				SetIdentifier(handler_seller.Alphabets[index]).
				SetName(item.Name).
				SetPrice(item.Price).
				SetStock(item.Stock),
		)
	}

	err = tx.PostItem.CreateBulk(postItems...).Exec(c.Context())
	if err != nil {
		log.Print(err)
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "Could not create post")
	}

	err = tx.Commit()
	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not commit transaction")
	}

	return c.JSON(fiber.Map{})
}

func GetPost(c *fiber.Ctx) error {
	postID := c.Params("id")

	p, err := database.DBConn.Post.
		Query().
		Select(
			post.FieldTitle, post.FieldBody,
			post.FieldStorageType, post.FieldDeadline, post.FieldDeliveryDate,
			post.FieldIsInStock, post.FieldImages,
		).
		Where(post.ID(postID)).
		WithSeller(func(uq *ent.UserQuery) {
			uq.Select(user.FieldDisplayName, user.FieldPictureURL)
		}).
		WithPostItems(func(piq *ent.PostItemQuery) {
			piq.Order(postitem.ByIdentifier())
		}).
		First(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Post not found")
	}

	return c.JSON(fiber.Map{
		"data": p,
	})
}

func GetPostFinanceDelivers(c *fiber.Ctx) error {
	postID := c.Params("postID")

	d, err := database.DBConn.Deliver.
		Query().
		Select(
			deliver.FieldNormalOrders, deliver.FieldExtraOrders,
			deliver.FieldNormalTotal, deliver.FieldNormalFee,
			deliver.FieldExtraTotal, deliver.FieldExtraFee,
			deliver.FieldCreatedAt,
		).
		Where(deliver.PostID(postID)).All(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query delivers")
	}

	return c.JSON(fiber.Map{
		"data": d,
	})
}

type updatePostStatusForm struct {
	Status post.Status `json:"status"`
}

func UpdatePostStatus(c *fiber.Ctx) error {
	postID := c.Params("id")

	postForm := new(updatePostStatusForm)

	if err := c.BodyParser(postForm); err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusBadRequest, "Could not parse order form")
	}

	if err := post.StatusValidator(postForm.Status); err != nil {
		return fiber.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(`Post status "%v" is not valid`, postForm.Status),
		)
	}

	err := database.DBConn.Post.
		UpdateOneID(postID).
		SetStatus(postForm.Status).
		Exec(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not update post status")
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
