package postRequests

import (
	"dmp-training/constants"
	"github.com/gofiber/fiber/v2"
)

type DeletePostRequest struct {
	Id int `json:"id"`
}

func ValidateDeletePostRequest(c *fiber.Ctx) (valid bool, msg string, req *DeletePostRequest) {
	req = new(DeletePostRequest)

	if err := c.BodyParser(req); err != nil {
		msg = constants.MessageBadRequest
		return
	}

	if req.Id == 0 {
		msg = "Id is required"
		return
	}

	valid = true
	return
}
