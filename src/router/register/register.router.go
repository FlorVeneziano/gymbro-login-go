package register

import (
	"fmt"

	"github.com/FlorVeneziano/gymbro-login-go/db/users"
	services "github.com/FlorVeneziano/gymbro-login-go/services/register"
	"github.com/FlorVeneziano/gymbro-login-go/types"
	"github.com/gofiber/fiber/v2"
)

type registerSuccess struct {
	types.Response
	Data string `json:"data"`
}

func Register(c *fiber.Ctx) error {

	var req *types.RegisterDTO

	err := c.BodyParser(&req)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(types.Response{
			Success: false,
			Code:    400,
			Message: err.Error(),
		})
	}

	usr := users.NewUserProvider(c)

	response := services.NewRegisterService(usr).Register(req.Email, req.Password)
	fmt.Println(response)

	return c.JSON(registerSuccess{
		Response: response,
		Data:     "Register data",
	})

}
