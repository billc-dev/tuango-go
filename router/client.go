package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent/user"
	"github.com/billc-dev/tuango-go/handler_client"
	"github.com/gofiber/fiber/v2"
)

func SetupClientRoutes(app *fiber.App) {
	v1 := app.Group("/api/client/v1")

	// PUT update all
	// PATCH partial modification

	user := v1.Group("/user")
	user.Get("/", ClientAuthenticated, handler_client.GetUser) // get user data
	// user.Post("/login")                // login
	user.Get("/likes", ClientAuthenticated, handler_client.GetLikes)                               // get liked posts
	user.Get("/orders", ClientAuthenticated, handler_client.GetOrders)                             // get orders => query status
	user.Get("/delivered-order-count", ClientAuthenticated, handler_client.GetDeliveredOrderCount) // get delivered order count
	// user.Get("/notification-count")    // get notification count
	user.Get("/rooms", ClientAuthenticated, handler_client.GetRooms)

	posts := v1.Group("/posts")
	posts.Get("/", handler_client.GetPosts) // get posts => query postNum, status, title, dates
	// posts.Get("/hot")                                                         // get hot posts
	posts.Get("/:id", handler_client.GetPost)                                 // get post
	posts.Get("/:id/comments", handler_client.GetPostComments)                // get post comments
	posts.Get("/:id/orders", handler_client.GetPostOrders)                    // get post orders
	posts.Post("/:id/like", ClientAuthenticated, handler_client.LikePost)     // like a post
	posts.Delete("/:id/like", ClientAuthenticated, handler_client.UnlikePost) // unlike a post

	orders := v1.Group("/orders")
	orders.Post("/", ClientAuthenticated, handler_client.CreateOrder)      // create order
	orders.Delete("/:id", ClientAuthenticated, handler_client.CancelOrder) // delete order if post status is open and order status is ordered
}

func ClientAuthenticated(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")
	token := strings.Replace(authorizationHeader, "Bearer ", "", 1)

	if len(token) == 0 {
		return fiber.NewError(http.StatusUnauthorized, "Bearer token missing")
	}

	decodedToken := "5fef56d3ec7ace8690f408c2" // TODO: fix

	u, err := database.DBConn.User.
		Query().
		Where(user.ID(decodedToken)).
		Select(user.FieldID, user.FieldDisplayName, user.FieldRole, user.FieldStatus).
		First(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusUnauthorized, "User not found")
	}

	if *u.Status == user.StatusBlocked {
		return fiber.NewError(http.StatusUnauthorized, "User unauthorized")
	}

	c.Locals("user", u)

	return c.Next()
}
