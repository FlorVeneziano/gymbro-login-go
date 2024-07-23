package register_services

import (
	"fmt"
	"time"

	users "github.com/FlorVeneziano/gymbro-login-go/db/users"
	"github.com/FlorVeneziano/gymbro-login-go/helpers"
	pswd "github.com/FlorVeneziano/gymbro-login-go/helpers"
	"github.com/FlorVeneziano/gymbro-login-go/providers/envs"
	"github.com/FlorVeneziano/gymbro-login-go/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterService struct {
	usr users.UserProviderInterface
}

func NewRegisterService(usr users.UserProviderInterface) *RegisterService {
	return &RegisterService{
		usr: usr,
	}
}

func (l *RegisterService) Register(email, password string) (types.RequestResponse, error) {
	env := envs.GetEnvs()

	// * Get user by email
	user, err := l.usr.GetUserByEmail(email)

	if user != nil {
		return types.RequestResponse{
			Response: types.Response{
				Success: false,
				Code:    400,
				Message: "User already exists",
			},
		}, err
	}

	pass, pswdErr := pswd.HashPassword(password)

	if pswdErr != nil {
		return types.RequestResponse{
			Response: types.Response{
				Success: false,
				Code:    401,
				Message: "Error hashing password",
			},
		}, err
	}

	newUsr, err := l.usr.CreateUser(&users.User{
		Id:           primitive.NewObjectID(),
		Email:        email,
		PasswordHash: pass,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})

	if err != nil {
		return types.RequestResponse{
			Response: types.Response{
				Success: false,
				Code:    401,
				Message: "Error creating user",
			},
		}, err
	}

	auth := helpers.JWTAuth{
		SecretKey: env.JWT_SECRET,
	}

	fmt.Println(newUsr.Id, newUsr.Email)
	token, err := auth.GenerateToken(newUsr.Id.Hex(), newUsr.Email)

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
			Message: "Register successful",
		},
		Data: token,
	}, nil
}
