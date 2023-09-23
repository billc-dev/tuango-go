package admin

import (
	"log"
	"net/http"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent/complete"
	"github.com/billc-dev/tuango-go/ent/predicate"
	"github.com/gofiber/fiber/v2"
)

func GetCompletes(c *fiber.Ctx) error {
	m := c.Queries()

	completeWhere := []predicate.Complete{}

	if userID := m["user_id"]; userID != "" {
		completeWhere = append(completeWhere, complete.UserID(userID))
	}

	if linePay := m["line_pay"]; linePay != "" {
		completeWhere = append(completeWhere, complete.LinePay(linePay == "true"))
	}

	if confirmed := m["confirmed"]; confirmed != "" {
		completeWhere = append(completeWhere, complete.Confirmed(confirmed == "true"))
	}

	limit := 20
	page := 0

	parsedPage, err := strconv.Atoi(m["page"])
	if err == nil {
		page = parsedPage
	}

	offset := page * limit

	completes, err := database.DBConn.Complete.Query().
		Where(completeWhere...).
		Order(complete.ByCreatedAt(sql.OrderDesc())).
		Limit(limit).
		Offset(offset).
		All(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query completes")
	}

	count, err := database.DBConn.Complete.Query().Where(completeWhere...).Count(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query complete count")
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"completes":   completes,
			"count":       count,
			"hasNextPage": len(completes) == limit,
		},
	})
}
