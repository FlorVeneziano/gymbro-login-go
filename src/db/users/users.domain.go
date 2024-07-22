package users

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserProviderInterface interface {
	CreateUser(user *User) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserById(id string) (*User, error)
	UpdateUser(id, field, newValue string) (*User, error)
}

type userProvider struct {
	c      *fiber.Ctx
	client *mongo.Client
	coll   *mongo.Collection
}

type User struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email        string             `json:"email"`
	PasswordHash string             `json:"passwordHash"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
