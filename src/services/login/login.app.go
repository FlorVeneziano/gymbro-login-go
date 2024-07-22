package login_services

import (
	users "github.com/FlorVeneziano/gymbro-login-go/db/users"
	pswd "github.com/FlorVeneziano/gymbro-login-go/helpers"
	"github.com/FlorVeneziano/gymbro-login-go/types"
)

type LoginService struct {
	usr users.UserProviderInterface
}

func NewLoginService(usr users.UserProviderInterface) *LoginService {
	return &LoginService{
		usr: usr,
	}
}

func (l *LoginService) Login(email, password string) types.Response {
	// * Get user by email
	user, err := l.usr.GetUserByEmail(email)

	if err != nil {
		return types.Response{
			Success: false,
			Code:    400,
			Message: "Bad request",
		}
	}

	// * Compare password

	pswdErr := pswd.ComparePasswords(user.PasswordHash, password)

	if pswdErr != nil {
		return types.Response{
			Success: false,
			Code:    401,
			Message: "Wrong username or password",
		}
	}

	return types.Response{
		Success: true,
		Code:    200,
		Message: "Login successful",
	}
}
