package postRequests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muflikhandimasd/golang-basic-clean/constants"
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
