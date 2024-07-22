package register_services

import (
	"time"

	users "github.com/FlorVeneziano/gymbro-login-go/db/users"
	pswd "github.com/FlorVeneziano/gymbro-login-go/helpers"
	"github.com/FlorVeneziano/gymbro-login-go/types"
)

type RegisterService struct {
	usr users.UserProviderInterface
}

func NewRegisterService(usr users.UserProviderInterface) *RegisterService {
	return &RegisterService{
		usr: usr,
	}
}

func (l *RegisterService) Register(email, password string) types.Response {
	// * Get user by email
	user, _ := l.usr.GetUserByEmail(email)

	if user != nil {
		return types.Response{
			Success: false,
			Code:    400,
			Message: "User already exists",
		}
	}

	pass, pswdErr := pswd.HashPassword(password)

	if pswdErr != nil {
		return types.Response{
			Success: false,
			Code:    401,
			Message: "Error hashing password",
		}
	}

	_, err := l.usr.CreateUser(&users.User{
		Email:        email,
		PasswordHash: pass,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})

	if err != nil {
		return types.Response{
			Success: false,
			Code:    401,
			Message: "Error creating user",
		}
	}

	return types.Response{
		Success: true,
		Code:    200,
		Message: "Register successful",
	}
}
