package userRequests

import (
	"dmp-training/constants"
	"github.com/gofiber/fiber/v2"
)

type RegisterRequest LoginRequest

func ValidateRegisterRequest(c *fiber.Ctx) (valid bool, msg string, req *RegisterRequest) {
	req = new(RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		msg = constants.MessageBadRequest
		return
	}
	if req.Username == "" {
		msg = "Username is required"
		return
	}
	if req.Password == "" {
		msg = "Password is required"
		return
	}

	valid = true
	return
}
