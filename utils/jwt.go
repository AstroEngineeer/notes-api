package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint64) (string, error) {
	key := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user-id":    userID,
			"expires-at": time.Now().Add(time.Hour).Unix(),
			"issued-at":  time.Now().Unix(),
			"auth":       true,
			"issuer":     "login",
		})
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("Error signing token: %v", err)
	}

	return signedToken, nil
}
