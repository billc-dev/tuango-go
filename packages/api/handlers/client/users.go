package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	"github.com/billc-dev/tuango-go/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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

const JWT_SECRET = "lkasjdfkljalskdfjaslkdfj"

func Login() error {
	userID := ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID":    userID,
			"ExpiresAt": time.Now().Add(10 * time.Minute),
		},
	)

	accessString, err := token.SignedString(JWT_SECRET)
	if err != nil {
		log.Printf("Failed to sign the token due to: %v", err)
		return err
	}
	log.Print(accessString)
	return nil
}

type lineLoginResponse struct {
	IDToken string `json:"id_token"`
}

type MyCustomClaims struct {
	LineID  string `json:"sub"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	jwt.StandardClaims
}

func LineLogin(c *fiber.Ctx) error {
	code := c.Params("code")
	m := c.Queries()
	// log.Print(m["redirect_uri"])

	payload := strings.NewReader(fmt.Sprintf(
		"grant_type=authorization_code&code=%s&client_id=%s&client_secret=%s&redirect_uri=%s",
		code, "1654947889", "7c5d36284f09fad398d14d5cde0dee10", m["redirect_uri"],
	),
	)
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(payload)
	// str := buf.String()
	// fmt.Println(str)
	req, err := http.NewRequest(http.MethodPost, "https://api.line.me/oauth2/v2.1/token", payload)
	if err != nil {
		return utils.Error(err, http.StatusInternalServerError, "Could not create request")
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return utils.Error(err, http.StatusBadRequest, "Could not get profile")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return utils.Error(err, http.StatusInternalServerError, "Could not read body")
	}

	data := lineLoginResponse{}
	json.Unmarshal(body, &data)

	token, _ := jwt.ParseWithClaims(data.IDToken, &MyCustomClaims{}, nil)

	if token == nil {
		return utils.Error(err, http.StatusInternalServerError, "Could not parse line id_token")
	}

	claims := token.Claims.(*MyCustomClaims)

	return c.JSON(utils.Result{
		Data: map[string]string{
			"LineID":  claims.LineID,
			"Name":    claims.Name,
			"Picture": claims.Picture,
		},
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
