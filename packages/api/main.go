package main

import (
	"log"
	"os"

	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/router"
	seedfuncs "github.com/billc-dev/tuango-go/seedFuncs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"github.com/swaggo/swag"
	"github.com/swaggo/swag/v2/gen"

	_ "github.com/billc-dev/tuango-go/docs"
	_ "github.com/lib/pq"
)

// Main
//
//	@title						Tuango API
//	@securityDefinitions.apikey	BearerToken
//	@in							header
//	@name						Authorization
//	@description				Bearer token
func main() {
	godotenv.Load()

	err := gen.New().Build(&gen.Config{
		SearchDir:           "./",
		MainAPIFile:         "main.go",
		PropNamingStrategy:  swag.CamelCase,
		OutputDir:           "./docs",
		OutputTypes:         []string{"yaml", "json"},
		ParseDependency:     false,
		ParseVendor:         false,
		GenerateOpenAPI3Doc: true,
	})

	if err != nil {
		log.Fatalf("Failed to generate swagger.json err: %v", err)
	}

	client := database.New()

	database.DevelopmentMigrate()

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return os.Getenv("ENVIRONMENT") == "development"
		},
	}))

	v1 := app.Group("/api/v1")

	// v1.Get("/swagger/*", swagger.HandlerDefault)

	v1.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		DeepLinking:              true,
		DefaultModelsExpandDepth: -1,
		DocExpansion:             "list",
		PersistAuthorization:     true,
	}))

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

	log.Fatal(app.Listen(":5010"))

	defer client.Close()
}
