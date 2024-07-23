package helpers

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTAuth is a struct that holds the JWT secret key
type JWTAuth struct {
	SecretKey string
}

// GenerateToken generates a new JWT token
func (auth *JWTAuth) GenerateToken(userID, userEmail string) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["email"] = userEmail
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate the token string
	tokenString, err := token.SignedString([]byte(auth.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies the authenticity of a JWT token
func (auth *JWTAuth) VerifyToken(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		// Return the secret key
		return []byte(auth.SecretKey), nil
	})

	// Check if the token is valid
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	// Extract the userID from the claims
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["userID"].(string)

	return userID, nil
}
