package users

import (
	"github.com/FlorVeneziano/gymbro-login-go/db"
	"github.com/gofiber/fiber/v2"
)

func NewUserProvider(c *fiber.Ctx) UserProviderInterface {
	client := db.GetDatabase()
	collection := db.GetUsersCollection()

	return &userProvider{
		c:      c,
		client: client,
		coll:   collection,
	}
}

func (u *userProvider) CreateUser(user *User) (*User, error) {
	var response *User

	_, err := u.coll.InsertOne(u.c.Context(), user)
	if err != nil {
		return response, err
	}

	response = user

	return response, nil
}

func (u *userProvider) GetUserByEmail(email string) (*User, error) {
	var response *User

	err := u.coll.FindOne(u.c.Context(), map[string]string{"email": email}).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (u *userProvider) GetUserById(id string) (*User, error) {
	var response *User

	err := u.coll.FindOne(u.c.Context(), map[string]string{"_id": id}).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (u *userProvider) UpdateUser(id, field, newValue string) (*User, error) {
	var response *User

	_, err := u.coll.UpdateOne(u.c.Context(), map[string]string{"_id": id}, map[string]interface{}{"$set": map[string]string{field: newValue}})
	if err != nil {
		return response, err
	}

	response, err = u.GetUserById(id)
	if err != nil {
		return response, err
	}

	return response, nil
}
