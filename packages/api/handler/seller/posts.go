package seller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent"
	"github.com/billc-dev/tuango-go/ent/order"
	"github.com/billc-dev/tuango-go/ent/post"
	"github.com/billc-dev/tuango-go/ent/predicate"
	"github.com/billc-dev/tuango-go/ent/schema"
	"github.com/gofiber/fiber/v2"
)

var Alphabets = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

func GetPosts(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*ent.User)

	if !ok {
		return fiber.NewError(http.StatusNotFound, "User not found")
	}

	m := c.Queries()

	postWhere := []predicate.Post{
		post.SellerID(u.ID),
	}
	u64, err := strconv.ParseInt(m["postNum"], 10, 64)
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
		postWhere = append(postWhere,
			post.StatusEQ(post.StatusClosed),
			post.HasPostOrdersWith(
				order.OrderNumNEQ(0),
				order.StatusEQ(order.StatusOrdered),
			),
		)
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

	deliveryDate := m["deliveryDate"]
	if deliveryDate != "" {
		postWhere = append(postWhere, post.DeliveryDateEQ(deliveryDate))
	}

	limit := 20

	page := 0
	u64, err = strconv.ParseInt(m["page"], 10, 64)
	if err == nil {
		page = int(u64)
	}

	offset := page * limit

	posts, err := database.DBConn.Debug().Post.
		Query().
		Select(post.FieldPostNum, post.FieldTitle, post.FieldOrderCount).
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

type createPostForm struct {
	Title        string           `json:"title"`
	Body         string           `json:"body"`
	Deadline     string           `json:"deadline"`
	DeliveryDate string           `json:"delivery_date"`
	StorageType  post.StorageType `json:"storage_type"`
	Images       []schema.Image   `json:"images"`
	Items        []struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Stock float64 `json:"stock"`
	} `json:"items"`
}

func CreatePost(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*ent.User)

	if !ok {
		return fiber.NewError(http.StatusNotFound, "User not found")
	}

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

	previousPost, err := database.DBConn.Post.Query().
		Where(post.StatusNEQ(post.StatusCanceled)).
		Select(post.FieldPostNum).
		Order(post.ByPostNum(sql.OrderDesc())).
		First(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query previous post")
	}

	newPost, err := tx.Post.Create().
		SetSellerID(u.ID).
		SetPostNum(*previousPost.PostNum + 1).
		SetTitle(postForm.Title).
		SetBody(postForm.Body).
		SetStorageType(postForm.StorageType).
		SetDeadline(postForm.Deadline).
		SetDeliveryDate(postForm.DeliveryDate).
		SetImages(postForm.Images).
		Save(c.Context())

	if err != nil {
		log.Print(err)
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "Could not create post")
	}

	if len(postForm.Items) > len(Alphabets) {
		tx.Rollback()
		return fiber.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf("Post items length is more than %v", len(Alphabets)),
		)
	}

	postItems := []*ent.PostItemCreate{}

	for index, item := range postForm.Items {
		postItems = append(postItems,
			tx.PostItem.Create().
				SetPostID(newPost.ID).
				SetIdentifier(Alphabets[index]).
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

	u, ok := c.Locals("user").(*ent.User)

	if !ok {
		return fiber.NewError(http.StatusNotFound, "User not found")
	}

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
		Where(post.SellerID(u.ID)).
		SetTitle(postForm.Title).
		SetBody(postForm.Body).
		SetStorageType(postForm.StorageType).
		SetDeadline(postForm.Deadline).
		SetDeliveryDate(postForm.DeliveryDate).
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
				SetIdentifier(Alphabets[index]).
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

func ClosePost(c *fiber.Ctx) error {
	postID := c.Params("id")

	u, ok := c.Locals("user").(*ent.User)

	if !ok {
		return fiber.NewError(http.StatusNotFound, "User not found")
	}

	err := database.DBConn.Debug().Post.
		UpdateOneID(postID).
		Where(post.SellerID(u.ID), post.StatusEQ(post.StatusOpen)).
		SetStatus(post.StatusClosed).
		Exec(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not close post")
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
