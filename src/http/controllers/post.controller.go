package controllers

import (
	"dmp-training/src/domain/usecases"
	"dmp-training/src/http/middlewares"
	"dmp-training/src/http/requests/postRequests"
	"dmp-training/src/http/responses"
	"github.com/gofiber/fiber/v2"
)

type PostController interface {
	Groups(group fiber.Router)
	GetAll(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type postController struct {
	uc usecases.PostUseCase
}

func NewPostController(uc usecases.PostUseCase) PostController {
	return &postController{uc: uc}
}

func (a *postController) Groups(group fiber.Router) {
	itemsGroup := group.Group("/v1/posts")
	middleware := middlewares.NewJWT
	itemsGroup.Get("/", middleware, a.GetAll)
	itemsGroup.Post("/", middleware, a.Create)
	itemsGroup.Post("/update", middleware, a.Update)
	itemsGroup.Post("/delete", middleware, a.Delete)

}

func (a *postController) GetAll(c *fiber.Ctx) error {
	valid, msg, req := postRequests.ValidateGetAllPostRequest(c)

	if !valid {
		return responses.ResponseErrorBadRequest(c, msg)
	}

	code, msg, resUseCase := a.uc.GetAll(c.Context(), req)

	return responses.ResponseConverter(code, c, resUseCase, msg)
}

func (a *postController) Create(c *fiber.Ctx) error {
	valid, msg, req := postRequests.ValidateCreatePostRequest(c)

	if !valid {
		return responses.ResponseErrorBadRequest(c, msg)
	}

	code, msg := a.uc.Create(c.Context(), req)

	return responses.ResponseConverter(code, c, nil, msg)
}

func (a *postController) Update(c *fiber.Ctx) error {
	valid, msg, req := postRequests.ValidateUpdatePostRequest(c)

	if !valid {
		return responses.ResponseErrorBadRequest(c, msg)
	}

	code, msg := a.uc.Update(c.Context(), req)

	return responses.ResponseConverter(code, c, nil, msg)
}

func (a *postController) Delete(c *fiber.Ctx) error {
	valid, msg, req := postRequests.ValidateDeletePostRequest(c)

	if !valid {
		return responses.ResponseErrorBadRequest(c, msg)
	}

	code, msg := a.uc.Delete(c.Context(), req)

	return responses.ResponseConverter(code, c, nil, msg)
}
