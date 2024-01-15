package postRequests

import (
	"dmp-training/constants"
	"dmp-training/helpers"
	"github.com/gofiber/fiber/v2"
)

type GetAllPostRequest struct {
	UserId int32 `json:"user_id"`
}

func ValidateGetAllPostRequest(c *fiber.Ctx) (valid bool, msg string, req *GetAllPostRequest) {
	req = new(GetAllPostRequest)
	claim, err := helpers.ParseToken(c)

	if err != nil {
		msg = constants.MessageForbidden
		return
	}

	req.UserId = claim.Id

	valid = true
	return
}
