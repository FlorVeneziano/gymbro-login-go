package login

import (
	"github.com/FlorVeneziano/gymbro-login-go/db/users"
	services "github.com/FlorVeneziano/gymbro-login-go/services/login"
	"github.com/FlorVeneziano/gymbro-login-go/types"
	"github.com/gofiber/fiber/v2"
)

type loginSuccess struct {
	types.Response
	Data string `json:"data"`
}

func Login(c *fiber.Ctx) error {

	var req types.LoginDTO

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(types.Response{
			Success: false,
			Code:    400,
			Message: "Bad request",
		})
	}

	usr := users.NewUserProvider(c)

	response := services.NewLoginService(usr).Login(req.Email, req.Password)

	return c.JSON(loginSuccess{
		Response: response,
		Data:     "Login data",
	})
}
