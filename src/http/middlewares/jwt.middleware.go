package middlewares

import (
	"dmp-training/constants"
	"dmp-training/helpers"
	"dmp-training/src/http/responses"
	"github.com/gofiber/fiber/v2"
)

func NewJWT(c *fiber.Ctx) error {
	if _, err := helpers.ParseToken(c); err != nil {
		return responses.ResponseErrorUnauthorized(c, constants.MessageForbidden)
	}
	return c.Next()
}
