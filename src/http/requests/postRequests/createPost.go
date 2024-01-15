package postRequests

import (
	"dmp-training/constants"
	"dmp-training/helpers"
	"github.com/gofiber/fiber/v2"
)

type CreatePostRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int32  `json:"user_id"`
}

func ValidateCreatePostRequest(c *fiber.Ctx) (valid bool, msg string, req *CreatePostRequest) {
	req = new(CreatePostRequest)

	claim, err := helpers.ParseToken(c)

	if err != nil {
		msg = constants.MessageForbidden
		return
	}

	if err = c.BodyParser(req); err != nil {
		msg = constants.MessageBadRequest
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

	req.UserId = claim.Id
	valid = true
	return
}
