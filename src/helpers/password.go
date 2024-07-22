package helpers

import (
	"github.com/FlorVeneziano/gymbro-login-go/providers/envs"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	pepperedPassword := password + envs.GetEnvs().PEPPER
	bytes, err := bcrypt.GenerateFromPassword([]byte(pepperedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ComparePasswords(hashedPassword, plainPassword string) error {
	pepperedPassword := plainPassword + envs.GetEnvs().PEPPER
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(pepperedPassword))
}
