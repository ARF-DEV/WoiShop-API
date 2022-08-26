package helpers

import (
	"azura-lab-intern/study-case-1/models"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	jwt.StandardClaims
	userData models.User
}

func GenerateUserToken(user models.User, expTime time.Time) (string, error) {
	key := os.Getenv("JWT_SECRET_KEY")

	claims := UserClaims{
		userData: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Println("Error while generating Token: ", err.Error())
		return "", err
	}

	return tokenString, nil
}
