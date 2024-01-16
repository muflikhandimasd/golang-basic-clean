package userRequests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muflikhandimasd/golang-basic-clean/constants"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ValidateLoginRequest(c *fiber.Ctx) (valid bool, msg string, req *LoginRequest) {
	req = new(LoginRequest)
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
