package router

import (
	routerLogin "github.com/FlorVeneziano/gymbro-login-go/router/login"
	routerRegister "github.com/FlorVeneziano/gymbro-login-go/router/register"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("Nothing to see here") })
	app.Get("/health", func(c *fiber.Ctx) error { return c.SendString("OK") })
	// * Setup routes

	app.Post("/login", routerLogin.Login)
	app.Post("/register", routerRegister.Register)
	// app.Post("/forgot-password", forgotPassword)
	// app.Post("/reset-password", resetPassword)
	// app.Post("/change-password", changePassword)
	// app.Post("/refresh-token", refreshToken)
	// app.Post("/logout", logout)
}
