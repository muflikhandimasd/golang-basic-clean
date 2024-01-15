package userRequests

import (
	"github.com/gofiber/fiber/v2"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
	Id           int32  `json:"id"`
	Username     string `json:"username"`
}

func ValidateRefreshTokenRequest(c *fiber.Ctx) (valid bool, msg string, req *RefreshTokenRequest) {
	req = new(RefreshTokenRequest)
	if err := c.BodyParser(req); err != nil {
		msg = "Refresh token is required"
		return
	}
	if req.RefreshToken == "" {
		msg = "Refresh token is required"
		return
	}

	valid = true
	return
}
