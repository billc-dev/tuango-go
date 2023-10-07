package router

import (
	"net/http"
	"strings"

	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent/user"
	"github.com/billc-dev/tuango-go/handlers/client"
	"github.com/billc-dev/tuango-go/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func SetupClientRoutes(app *fiber.App) {
	v1 := app.Group("/api/client/v1")

	// PUT update all
	// PATCH partial modification

	user := v1.Group("/user")
	user.Get("/", ClientAuthenticated, client.GetUser) // get user data
	// user.Post("/refresh") // refresh jwt token
	user.Get("/likes", ClientAuthenticated, client.GetLikes)                               // get liked posts
	user.Get("/orders", ClientAuthenticated, client.GetOrders)                             // get orders => query status
	user.Get("/delivered-order-count", ClientAuthenticated, client.GetDeliveredOrderCount) // get delivered order count
	// user.Get("/notification-count")    // get notification count
	user.Get("/rooms", ClientAuthenticated, client.GetRooms)
	user.Post("/login/line/:code", client.LineLogin)

	posts := v1.Group("/posts")
	posts.Get("/", client.GetPosts) // get posts => query postNum, status, title, dates
	// posts.Get("/hot")                                                         // get hot posts
	posts.Get("/:id", client.GetPost)                                 // get post
	posts.Get("/:id/comments", client.GetPostComments)                // get post comments
	posts.Get("/:id/orders", client.GetPostOrders)                    // get post orders
	posts.Post("/:id/like", ClientAuthenticated, client.LikePost)     // like a post
	posts.Delete("/:id/like", ClientAuthenticated, client.UnlikePost) // unlike a post

	orders := v1.Group("/orders")
	orders.Post("/", ClientAuthenticated, client.CreateOrder)      // create order
	orders.Delete("/:id", ClientAuthenticated, client.CancelOrder) // delete order if post status is open and order status is ordered
}

func ClientAuthenticated(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")
	token := strings.Replace(authorizationHeader, "Bearer ", "", 1)

	if len(token) == 0 {
		return utils.Error(nil, http.StatusUnauthorized, "Bearer token missing")
	}

	type MyClaims struct {
		jwt.StandardClaims
	}

	parsedToken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return client.JWT_SECRET, nil
	})

	if err != nil {
		return utils.Error(err, http.StatusBadRequest, "Could not parse token")
	}

	claims := parsedToken.Claims.(*MyClaims)

	u, err := database.DBConn.User.
		Query().
		Where(user.ID(claims.Id)).
		Select(user.FieldID, user.FieldDisplayName, user.FieldRole, user.FieldStatus).
		First(c.Context())

	if err != nil {
		return utils.Error(err, http.StatusUnauthorized, "User not found")
	}

	if *u.Status == user.StatusBlocked {
		return utils.Error(nil, http.StatusUnauthorized, "User unauthorized")
	}

	c.Locals("user", u)

	return c.Next()
}
