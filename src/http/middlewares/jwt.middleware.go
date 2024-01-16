package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muflikhandimasd/golang-basic-clean/constants"
	"github.com/muflikhandimasd/golang-basic-clean/helpers"
	"github.com/muflikhandimasd/golang-basic-clean/src/http/responses"
)

func NewJWT(c *fiber.Ctx) error {
	if _, err := helpers.ParseToken(c); err != nil {
		return responses.ResponseErrorUnauthorized(c, constants.MessageForbidden)
	}
	return c.Next()
}
