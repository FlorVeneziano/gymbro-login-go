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
		return c.Status(fiber.StatusBadRequest).JSON(types.RequestResponse{
			Response: types.Response{
				Success: false,
				Code:    400,
				Message: "Bad request",
			},
			Data: err.Error(),
		})
	}

	usr := users.NewUserProvider(c)

	response, errLogin := services.NewLoginService(usr).Login(req.Email, req.Password)

	if errLogin != nil {
		return c.Status(fiber.StatusBadRequest).JSON(types.RequestResponse{
			Response: types.Response{
				Success: false,
				Code:    400,
				Message: "Bad request",
			},
			Data: errLogin.Error(),
		})
	}

	return c.JSON(loginSuccess{
		Response: response.Response,
		Data:     response.Data,
	})
}
