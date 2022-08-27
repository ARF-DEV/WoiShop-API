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
	UserData models.User `json:"user_data"`
}

type RegisterOTPClaims struct {
	jwt.StandardClaims
	UserData  models.User `json:"user_data"`
	OTPSecret string
}

func GenerateRegisterOTPClaims(user models.User, expTime time.Time) (string, string, error) {
	jwtKey := os.Getenv("JWT_OTP_KEY")

	key, err := GenerateOTPcode(6)

	if err != nil {
		return "", "", err
	}

	claims := RegisterOTPClaims{
		UserData:  user,
		OTPSecret: key,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		return "", "", err
	}

	return tokenString, key, nil
}

func GenerateUserToken(user models.User, expTime time.Time) (string, error) {
	key := os.Getenv("JWT_SECRET_KEY")

	claims := UserClaims{
		UserData: user,
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
