package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent/user"
	"github.com/billc-dev/tuango-go/handlers/admin"
	"github.com/gofiber/fiber/v2"
)

func SetupAdminRoutes(app *fiber.App) {
	v1 := app.Group("/api/admin/v1")

	// Add sorting?

	posts := v1.Group("/posts")
	posts.Get("/", AdminAuthenticated, admin.GetPosts)
	posts.Post("/", AdminAuthenticated, admin.CreatePost)
	posts.Get("/:id", AdminAuthenticated, admin.GetPost)
	posts.Put("/:id", AdminAuthenticated, admin.UpdatePost)                               // update post
	posts.Patch("/:id/deliver/finance", AdminAuthenticated, admin.GetPostFinanceDelivers) // get post finance
	// posts.Patch("/:id/deliver/finance") // update post finance
	posts.Patch("/:id/status", AdminAuthenticated, admin.UpdatePostStatus) // update post status
	posts.Get("/date/:date", AdminAuthenticated, admin.GetPostsByDate)

	orders := v1.Group("/orders")
	orders.Get("/", AdminAuthenticated, admin.GetOrders) // get orders => query pickupNum, postNum, text, status, is-in-stock
	// orders.Get("/location", AdminAuthenticated)                     // get orders by location
	orders.Get("/:id", AdminAuthenticated, admin.GetOrder)  // get order
	orders.Post("/", AdminAuthenticated, admin.CreateOrder) // create order
	// orders.Put("/:id") // update order
	// orders.Post("/extra", AdminAuthenticated, admin.CreateExtraOrder) // create extra order

	users := v1.Group("/users")
	users.Get("/", AdminAuthenticated, admin.GetUsers)   // get users => query pickupNum, userId, notified, status, role
	users.Get("/:id", AdminAuthenticated, admin.GetUser) // get user
	// users.Put("/:id")           // update user
	users.Patch("/:id/approve", AdminAuthenticated, admin.ApproveUser) // approve user

	delivers := v1.Group("/delivers")
	delivers.Get("/", AdminAuthenticated, admin.GetDelivers) // get delivers

	completes := v1.Group("/completes")
	completes.Get("/", AdminAuthenticated, admin.GetCompletes) // get completes => query userId, line_pay, confirmed

	notifications := v1.Group("/notifications")
	notifications.Get("/", AdminAuthenticated, admin.GetNotifications) // get notifications

	report := v1.Group("/report")
	report.Get("/finance", admin.GetFinance) // sum total
}

func AdminAuthenticated(c *fiber.Ctx) error {
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

	if *u.Role != user.RoleAdmin {
		return fiber.NewError(http.StatusUnauthorized, "User unauthorized")
	}

	c.Locals("user", u)

	return c.Next()
}
