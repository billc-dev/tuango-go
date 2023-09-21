package main

import (
	"log"

	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent"
	"github.com/billc-dev/tuango-go/ent/complete"
	"github.com/billc-dev/tuango-go/ent/order"
	"github.com/billc-dev/tuango-go/ent/post"
	"github.com/billc-dev/tuango-go/ent/user"
	"github.com/billc-dev/tuango-go/router"
	seedfuncs "github.com/billc-dev/tuango-go/seedFuncs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/lib/pq"
)

func main() {
	client := database.New()

	database.DevelopmentMigrate()

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	v1 := app.Group("/api/v1")

	seed := v1.Group("/seed")
	seed.Get("/all", seedfuncs.SeedAll(client))
	seed.Get("/users", seedfuncs.SeedUsers(client))         // 1
	seed.Get("/posts", seedfuncs.SeedPosts(client))         // 2
	seed.Get("/orders", seedfuncs.SeedOrders(client))       // 3
	seed.Get("/completes", seedfuncs.SeedCompletes(client)) // 4
	seed.Get("/comments", seedfuncs.SeedComments(client))   // 5
	seed.Get("/notifies", seedfuncs.SeedNotifies(client))   // 6
	seed.Get("/delivers", seedfuncs.SeedDelivers(client))   // 7
	seed.Get("/likes", seedfuncs.SeedLikes(client))         // 8
	seed.Get("/rooms", seedfuncs.SeedRooms(client))         // 9
	seed.Get("/messages", seedfuncs.SeedMessages(client))   // 10

	router.SetupClientRoutes(app)
	router.SetupSellerRoutes(app)
	router.SetupAdminRoutes(app)

	v1.Get("/posts", func(c *fiber.Ctx) error {
		posts, err := client.Post.
			Query().
			WithSeller(func(uq *ent.UserQuery) {
				uq.Select(user.FieldDisplayName)
			}).
			Limit(10).
			Order().
			Select().
			All(c.Context())

		if err != nil {
			return err
		}

		return c.JSON(posts)
	})

	v1.Get("/posts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		post, err := client.Post.
			Query().
			Where(post.ID(id), post.StatusNEQ(post.StatusCanceled)).
			WithPostItems().
			First(c.Context())
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		log.Print(post.Edges.PostItems)
		return c.JSON(post.Edges.PostItems)
	})

	v1.Get("/users", func(c *fiber.Ctx) error {
		users, err := client.User.
			Query().
			Order(user.ByID(sql.OrderDesc())).
			Limit(3).
			All(c.Context())

		if err != nil {
			return err
		}

		return c.JSON(users)
	})

	v1.Get("/orders", func(c *fiber.Ctx) error {
		users, err := client.Order.
			Query().
			Where(order.HasPostWith(post.ID("64492b8b0424647496389d32"))).
			WithPost(func(pq *ent.PostQuery) {
				pq.Select(post.FieldPostNum, post.FieldTitle)
			}).
			WithUser(func(uq *ent.UserQuery) {
				uq.Select(user.FieldDisplayName, user.FieldPictureURL)
			}).
			Order(order.ByOrderNum(sql.OrderDesc())).
			Limit(3).
			All(c.Context())

		if err != nil {
			return err
		}

		return c.JSON(users)
	})

	type CreateUserData struct {
		Username    string    `json:"username"`
		DisplayName string    `json:"displayName"`
		PictureUrl  string    `json:"pictureUrl"`
		Role        user.Role `json:"role"`
	}

	v1.Post("/users", func(c *fiber.Ctx) error {
		var d CreateUserData

		if err := c.BodyParser(&d); err != nil {
			return err
		}

		u, _ := client.User.Query().First(c.Context())

		return c.JSON(u)
	})

	v1.Get("/user-orders", func(c *fiber.Ctx) error {
		orders, err := client.Order.Query().
			Where(order.UserID("5fee9234848c444b0c8fed78")).
			Order(order.ByCreatedAt(sql.OrderDesc())).
			WithOrderItems().
			WithPost(func(pq *ent.PostQuery) {
				pq.Select(post.FieldPostNum, post.FieldTitle).WithSeller(func(uq *ent.UserQuery) {
					uq.Select(user.FieldDisplayName)
				})
			}).
			Limit(10).
			All(c.Context())

		if err != nil {
			return err
		}

		return c.JSON(orders)
	})

	v1.Get("/user-completes", func(c *fiber.Ctx) error {
		completes, err := client.Complete.Query().
			Where(complete.UserID("5fee9234848c444b0c8fed78")).
			Order(complete.ByCreatedAt(sql.OrderDesc())).
			Limit(10).
			All(c.Context())

		if err != nil {
			return err
		}

		return c.JSON(completes)
	})

	log.Fatal(app.Listen(":5010"))

	defer client.Close()
}
