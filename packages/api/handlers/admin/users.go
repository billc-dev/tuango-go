package admin

import (
	"log"
	"net/http"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent/predicate"
	"github.com/billc-dev/tuango-go/ent/user"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	m := c.Queries()

	userWhere := []predicate.User{}

	if pickupNum := m["pickup_num"]; pickupNum != "" {
		pickupNum, err := strconv.Atoi(pickupNum)
		if err == nil {
			userWhere = append(userWhere, user.PickupNum(float64(pickupNum)))
		}
	}

	if displayName := m["display_name"]; displayName != "" {
		userWhere = append(userWhere, user.DisplayNameContains(displayName))
	}

	if notified := m["notified"]; notified != "" {
		userWhere = append(userWhere, user.Notified(notified == "true"))
	}

	if status := m["status"]; status != "" {
		status := user.Status(status)
		userWhere = append(userWhere, user.StatusEQ(status))
	}

	if role := m["role"]; role != "" {
		role := user.Role(role)
		userWhere = append(userWhere, user.RoleEQ(role))
	}

	limit := 20
	page := 0

	parsedPage, err := strconv.Atoi(m["page"])
	if err == nil {
		page = parsedPage
	}

	offset := page * limit

	posts, err := database.DBConn.User.Query().
		Where(userWhere...).
		Order(user.ByPickupNum(sql.OrderDesc())).
		Limit(limit).
		Offset(offset).
		All(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query users")
	}

	count, err := database.DBConn.User.Query().Where(userWhere...).Count(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query user count")
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"users":       posts,
			"count":       count,
			"hasNextPage": len(posts) == limit,
		},
	})
}

func GetUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	u, err := database.DBConn.User.
		Query().
		Where(user.ID(userID)).
		First(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "User not found")
	}

	return c.JSON(fiber.Map{
		"data": u,
	})
}

func ApproveUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	err := database.DBConn.User.
		UpdateOneID(userID).
		SetStatus(user.StatusApproved).
		Exec(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not approve user")
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
