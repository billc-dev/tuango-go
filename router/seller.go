package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent/user"
	"github.com/billc-dev/tuango-go/handler_seller"
	"github.com/gofiber/fiber/v2"
)

func SetupSellerRoutes(app *fiber.App) {
	v1 := app.Group("/api/seller/v1")

	posts := v1.Group("/posts")
	posts.Get("/", SellerAuthenticated, handler_seller.GetPosts)             // get seller posts => query status
	posts.Post("/", SellerAuthenticated, handler_seller.CreatePost)          // create post
	posts.Put("/:id", SellerAuthenticated, handler_seller.UpdatePost)        // update post
	posts.Patch("/:id/close", SellerAuthenticated, handler_seller.ClosePost) // update post status to closed

	// orders := v1.Group("/orders")
	// orders.Patch("/:id", SellerAuthenticated, handler_seller.SetHasName)
}

func SellerAuthenticated(c *fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")
	token := strings.Replace(authorizationHeader, "Bearer ", "", 1)

	if len(token) == 0 {
		return fiber.NewError(http.StatusUnauthorized, "Bearer token missing")
	}

	decodedToken := "603dc6aacd70a28569ec2c59" // TODO: fix

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

	if *u.Role != user.RoleAdmin && *u.Role != user.RoleSeller {
		return fiber.NewError(http.StatusUnauthorized, "User unauthorized")
	}

	c.Locals("user", u)

	return c.Next()
}
