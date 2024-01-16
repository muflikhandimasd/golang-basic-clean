package postRequests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muflikhandimasd/golang-basic-clean/constants"
	"github.com/muflikhandimasd/golang-basic-clean/helpers"
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
