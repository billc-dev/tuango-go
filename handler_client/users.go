package handler_client

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent"
	"github.com/billc-dev/tuango-go/ent/complete"
	"github.com/billc-dev/tuango-go/ent/like"
	"github.com/billc-dev/tuango-go/ent/order"
	"github.com/billc-dev/tuango-go/ent/post"
	"github.com/billc-dev/tuango-go/ent/postitem"
	"github.com/billc-dev/tuango-go/ent/predicate"
	"github.com/billc-dev/tuango-go/ent/room"
	"github.com/billc-dev/tuango-go/ent/roomuser"
	"github.com/billc-dev/tuango-go/ent/user"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*ent.User)

	if !ok {
		return fiber.NewError(http.StatusNotFound, "User not found")
	}

	u, err := database.DBConn.User.Query().Where(user.ID(u.ID)).First(c.Context())

	if err != nil {
		return fiber.NewError(http.StatusNotFound, "User not found")
	}

	return c.JSON(fiber.Map{
		"data": u,
	})
}

func GetLikes(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*ent.User)

	if !ok {
		return fiber.NewError(http.StatusNotFound, "User not found")
	}

	likes, err := database.DBConn.Like.Query().Where(like.UserID(u.ID)).All(c.Context())

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Could not query likes")
	}

	return c.JSON(fiber.Map{
		"data": likes,
	})
}

func GetOrders(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*ent.User)

	if !ok {
		return fiber.NewError(http.StatusNotFound, "User not found")
	}

	m := c.Queries()

	status := order.Status(m["status"])
	err := order.StatusValidator(status)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, fmt.Sprintf(`Order status "%v" is not valid`, status))
	}

	text := m["text"]

	limit := 20

	page := 0
	u64, err := strconv.ParseUint(m["page"], 10, 64)
	if err == nil {
		page = int(u64)
	}

	offset := page * limit

	if status == order.StatusCompleted {
		completeWhere := []predicate.Complete{
			complete.UserID(u.ID),
		}
		if text != "" {
			completeWhere = append(completeWhere,
				func(s *sql.Selector) {
					s.Where(sql.ExprP("orders::text ILIKE $2", "%"+text+"%"))
				},
			)
		}

		orders, err := database.DBConn.Debug().Complete.
			Query().
			Order(complete.ByCreatedAt(sql.OrderDesc())).
			Where(completeWhere...).
			Limit(limit).
			Offset(offset).
			All(c.Context())

		if err != nil {
			log.Print(err)
			return fiber.NewError(http.StatusInternalServerError, "Could not query completed orders")
		}

		count, err := database.DBConn.Complete.Query().Where(completeWhere...).Count(c.Context())

		if err != nil {
			log.Print(err)
			return fiber.NewError(http.StatusInternalServerError, "Could not query completed order count")
		}

		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"orders":      orders,
				"count":       count,
				"hasNextPage": len(orders) == limit,
			},
		})

	}

	orderWhere := []predicate.Order{
		order.UserID(u.ID),
		order.StatusEQ(status),
	}

	if text != "" {
		orderWhere = append(orderWhere,
			order.HasPostWith(
				post.Or(
					post.TitleContains(text),
					post.BodyContains(text),
					post.HasPostItemsWith(
						postitem.NameContains(text),
					),
				),
			))
	}

	orders, err := database.DBConn.Debug().Order.
		Query().
		Order(order.ByCreatedAt(sql.OrderDesc())).
		Where(orderWhere...).
		WithOrderItems().
		Limit(limit).
		Offset(offset).
		All(c.Context())

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Could not query orders")
	}

	count, err := database.DBConn.Order.Query().Where(orderWhere...).Count(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query order count")
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"orders":      orders,
			"count":       count,
			"hasNextPage": len(orders) == limit,
		},
	})
}

func GetDeliveredOrderCount(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*ent.User)

	if !ok {
		return fiber.NewError(http.StatusNotFound, "User not found")
	}

	count, err := database.DBConn.Order.
		Query().
		Where(order.UserID(u.ID), order.StatusEQ(order.StatusDelivered)).
		Count(c.Context())

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Could not query order count")
	}

	return c.JSON(fiber.Map{
		"data": count,
	})
}

func GetRooms(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*ent.User)

	if !ok {
		return fiber.NewError(http.StatusNotFound, "User not found")
	}

	limit := 20

	roomWhere := []predicate.Room{
		room.HasRoomUsersWith(roomuser.UserID(u.ID)),
	}

	rooms, err := database.DBConn.Debug().Room.
		Query().
		Where(roomWhere...).
		WithRoomUsers(func(ruq *ent.RoomUserQuery) {
			ruq.Where(roomuser.UserIDNEQ(u.ID)).
				WithUser(func(uq *ent.UserQuery) {
					uq.Select(user.FieldDisplayName)
				})
		}).
		Order(room.ByUpdatedAt(sql.OrderDesc())).
		Limit(limit).
		All(c.Context())

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Could not query rooms")
	}

	count, err := database.DBConn.Room.
		Query().
		Where(roomWhere...).
		Count(c.Context())

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Could not query room count")
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"rooms":       rooms,
			"count":       count,
			"hasNextPage": len(rooms) == limit,
		},
	})
}

func limitRows(partitionBy string, limit int, orderBy ...string) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		if len(orderBy) == 0 {
			orderBy = append(orderBy, "id")
		}
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderBy(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}
