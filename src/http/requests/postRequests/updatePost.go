package postRequests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muflikhandimasd/golang-basic-clean/constants"
)

type UpdatePostRequest struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func ValidateUpdatePostRequest(c *fiber.Ctx) (valid bool, msg string, req *UpdatePostRequest) {
	req = new(UpdatePostRequest)

	if err := c.BodyParser(req); err != nil {
		msg = constants.MessageBadRequest
		return
	}

	if req.Id == 0 {
		msg = "Id is required"
		return
	}
	if req.Title == "" {
		msg = "Title is required"
		return
	}

	if req.Body == "" {
		msg = "Body is required"
		return
	}

	valid = true
	return
}
