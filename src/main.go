package main

import (
	"fmt"

	"github.com/FlorVeneziano/gymbro-login-go/db"
	"github.com/FlorVeneziano/gymbro-login-go/providers/envs"
	"github.com/FlorVeneziano/gymbro-login-go/router"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func main() {
	// * Get envs
	env := envs.GetEnvs()

	port := env.PORT

	db.GetDatabase()
	defer db.DisconnectDatabase()

	app := fiber.New(fiber.Config{
		ErrorHandler:          errorHandler,
		JSONEncoder:           sonic.Marshal,
		JSONDecoder:           sonic.Unmarshal,
		DisableStartupMessage: false,
	})

	app.Use(cors.New())
	app.Use(healthcheck.New())

	// * Setup routes
	router.SetupRoutes(app)

	fmt.Printf("Listening on port: %s", port)
	err := app.Listen(":" + port)
	if err != nil {
		fmt.Printf("Error starting server: %s", err)
	}

}

func errorHandler(c *fiber.Ctx, err error) error {
	// Custom error handling logic here
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Internal Server Error",
	})
}
