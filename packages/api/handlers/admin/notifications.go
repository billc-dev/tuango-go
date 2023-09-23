package admin

import (
	"log"
	"net/http"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent/notification"
	"github.com/gofiber/fiber/v2"
)

func GetNotifications(c *fiber.Ctx) error {
	m := c.Queries()

	limit := 20
	page := 0

	parsedPage, err := strconv.Atoi(m["page"])
	if err == nil {
		page = parsedPage
	}

	offset := page * limit

	notifications, err := database.DBConn.Notification.Query().
		Order(notification.ByCreatedAt(sql.OrderDesc())).
		Limit(limit).
		Offset(offset).
		All(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query notifications")
	}

	count, err := database.DBConn.Notification.Query().Count(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query notification count")
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"notifications": notifications,
			"count":         count,
			"hasNextPage":   len(notifications) == limit,
		},
	})
}
