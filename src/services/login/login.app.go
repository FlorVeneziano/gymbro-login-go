package login_services

import (
	users "github.com/FlorVeneziano/gymbro-login-go/db/users"
	helpers "github.com/FlorVeneziano/gymbro-login-go/helpers"
	"github.com/FlorVeneziano/gymbro-login-go/providers/envs"
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

func (l *LoginService) Login(email, password string) (types.RequestResponse, error) {
	env := envs.GetEnvs()
	// * Get user by email
	user, err := l.usr.GetUserByEmail(email)

	if err != nil {
		return types.RequestResponse{
			Response: types.Response{
				Success: false,
				Code:    400,
				Message: "Bad request",
			},
		}, err
	}

	// * Compare password

	pswdErr := helpers.ComparePasswords(user.PasswordHash, password)

	if pswdErr != nil {
		return types.RequestResponse{
			Response: types.Response{
				Success: false,
				Code:    401,
				Message: "Wrong username or password",
			},
		}, pswdErr
	}

	auth := helpers.JWTAuth{
		SecretKey: env.JWT_SECRET,
	}
	// * Generate JWT token
	token, err := auth.GenerateToken(user.Id.Hex(), user.Email)

	if err != nil {
		return types.RequestResponse{
			Response: types.Response{

				Success: false,
				Code:    500,
				Message: "Internal server error",
			},
		}, err
	}

	return types.RequestResponse{
		Response: types.Response{
			Success: true,
			Code:    200,
			Message: "Login successful",
		},
		Data: token,
	}, nil
}
