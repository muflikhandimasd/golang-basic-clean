package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muflikhandimasd/golang-basic-clean/src/domain/usecases"
	userRequests "github.com/muflikhandimasd/golang-basic-clean/src/http/requests/userRequests"
	"github.com/muflikhandimasd/golang-basic-clean/src/http/responses"
)

type AuthController interface {
	Groups(group fiber.Router)
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}

type authController struct {
	uc usecases.UserUseCase
}

func NewAuthController(uc usecases.UserUseCase) AuthController {
	return &authController{uc}
}

func (a *authController) Groups(group fiber.Router) {
	itemsGroup := group.Group("/v1/auth")
	itemsGroup.Post("/login", a.Login)
	itemsGroup.Post("/register", a.Register)
	itemsGroup.Post("/refresh-token", a.RefreshToken)
}

func (a *authController) Login(c *fiber.Ctx) error {
	valid, msg, req := userRequests.ValidateLoginRequest(c)

	if !valid {
		return responses.ResponseErrorBadRequest(c, msg)
	}

	code, msg, resUseCase := a.uc.Login(c.Context(), req)

	return responses.ResponseConverter(code, c, resUseCase, msg)
}

func (a *authController) Register(c *fiber.Ctx) error {
	valid, msg, req := userRequests.ValidateRegisterRequest(c)

	if !valid {
		return responses.ResponseErrorBadRequest(c, msg)
	}

	code, msg, resUseCase := a.uc.Register(c.Context(), req)

	return responses.ResponseConverter(code, c, resUseCase, msg)
}

func (a *authController) RefreshToken(c *fiber.Ctx) error {
	valid, msg, req := userRequests.ValidateRefreshTokenRequest(c)

	if !valid {
		return responses.ResponseErrorBadRequest(c, msg)
	}

	code, msg, resUseCase := a.uc.RefreshToken(c.Context(), req)

	return responses.ResponseConverter(code, c, resUseCase, msg)
}
