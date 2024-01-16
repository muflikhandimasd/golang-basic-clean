package helpers

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/muflikhandimasd/golang-basic-clean/constants"
)

type ClaimJWT struct {
	Id       int32  `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(id int32, username string) string {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	var key = []byte(os.Getenv("JWT_SECRET"))

	claims := ClaimJWT{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ParseInt64(os.Getenv("JWT_EXPIRES_IN"))) * time.Second)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	return signedToken
}

func GenerateRefreshToken(id int32, username string) string {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	var key = []byte(os.Getenv("JWT_REFRESH_SECRET"))

	claims := ClaimJWT{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ParseInt64(os.Getenv("JWT_REFRESH_EXPIRES_IN"))) * time.Second)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	return signedToken
}

func ParseToken(c *fiber.Ctx) (claims *ClaimJWT, err error) {
	if err = godotenv.Load(); err != nil {
		panic(err)
	}

	tokenString := c.Get("Authorization", "")
	if strings.TrimSpace(tokenString) == "" {
		err = errors.New(constants.MessageForbidden)
		return
	}

	tokenSplit := strings.Split(tokenString, " ")

	if len(tokenSplit) < 2 {
		err = errors.New(constants.MessageForbidden)
		return
	}

	if tokenSplit[0] != "Bearer" {
		err = errors.New(constants.MessageForbidden)
		return
	}

	realToken := tokenSplit[1]

	var key = []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(realToken, &ClaimJWT{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		err = errors.New(constants.MessageForbidden)
		return
	}

	claims, ok := token.Claims.(*ClaimJWT)
	if !ok {
		err = errors.New(constants.MessageForbidden)
		return
	}

	if !token.Valid {
		err = errors.New(constants.MessageForbidden)
		return
	}

	return

}

func ParseRefreshToken(resfreshToken string) (claims *ClaimJWT, err error) {
	if err = godotenv.Load(); err != nil {
		panic(err)
	}

	var key = []byte(os.Getenv("JWT_REFRESH_SECRET"))

	token, err := jwt.ParseWithClaims(resfreshToken, &ClaimJWT{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		err = errors.New(constants.MessageForbidden)
		return
	}

	claims, ok := token.Claims.(*ClaimJWT)
	if !ok {
		err = errors.New(constants.MessageForbidden)
		return
	}

	if !token.Valid {
		err = errors.New(constants.MessageForbidden)
		return
	}

	return

}
