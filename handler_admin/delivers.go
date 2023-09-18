package handler_admin

import (
	"log"
	"net/http"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent/deliver"
	"github.com/billc-dev/tuango-go/ent/post"
	"github.com/billc-dev/tuango-go/ent/predicate"
	"github.com/billc-dev/tuango-go/ent/user"
	"github.com/gofiber/fiber/v2"
)

func GetDelivers(c *fiber.Ctx) error {
	m := c.Queries()

	deliverWhere := []predicate.Deliver{}

	if userID := m["user_id"]; userID != "" {
		deliverWhere = append(deliverWhere, deliver.UserID(userID))
	}

	if displayName := m["display_name"]; displayName != "" {
		deliverWhere = append(deliverWhere, deliver.HasUserWith(user.DisplayNameContains(displayName)))
	}

	if text := m["text"]; text != "" {
		deliverWhere = append(deliverWhere, deliver.HasPostWith(
			post.Or(
				post.TitleContains(text),
				post.BodyContains(text),
			),
		))
	}

	limit := 20
	page := 0

	parsedPage, err := strconv.Atoi(m["page"])
	if err == nil {
		page = parsedPage
	}

	offset := page * limit

	delivers, err := database.DBConn.Deliver.Query().
		Where(deliverWhere...).
		Order(deliver.ByCreatedAt(sql.OrderDesc())).
		Limit(limit).
		Offset(offset).
		All(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query delivers")
	}

	count, err := database.DBConn.Deliver.Query().Where(deliverWhere...).Count(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query deliver count")
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"delivers":    delivers,
			"count":       count,
			"hasNextPage": len(delivers) == limit,
		},
	})
}
