package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent/user"
	"github.com/billc-dev/tuango-go/handlers/client"
	"github.com/gofiber/fiber/v2"
)

func SetupClientRoutes(app *fiber.App) {
	v1 := app.Group("/api/client/v1")

	// PUT update all
	// PATCH partial modification

	user := v1.Group("/user")
	user.Get("/", ClientAuthenticated, client.GetUser) // get user data
	// user.Post("/login")                // line login and generate jwt token
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
