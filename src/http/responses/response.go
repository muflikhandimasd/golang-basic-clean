package responses

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muflikhandimasd/golang-basic-clean/constants"
)

func ResponseConverter(code int, c *fiber.Ctx, data interface{}, msg string) error {
	if code == 200 {
		return ResponseSuccess(c, data)
	}
	switch code {
	case fiber.StatusBadRequest:
		return ResponseErrorBadRequest(c, msg)

	case fiber.StatusUnauthorized:
		return ResponseErrorUnauthorized(c, msg)

	case fiber.StatusForbidden:
		return ResponseErrorForbidden(c, msg)

	case fiber.StatusNotFound:
		return ResponseErrorNotFound(c, msg)

	case fiber.StatusInternalServerError:
		return ResponseErrorInternalServerError(c, msg)

	case fiber.StatusConflict:
		return ResponseErrorConflict(c, msg)

	default:
		return ResponseErrorInternalServerError(c, constants.MessageInternalServerError)
	}
}

func ResponseSuccess(c *fiber.Ctx, data interface{}) error {
	result := fiber.Map{
		"code":    fiber.StatusOK,
		"data":    data,
		"message": constants.MessageSuccess,
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func responseError(c *fiber.Ctx, code int, message string) error {
	result := fiber.Map{
		"code":    code,
		"data":    nil,
		"message": message,
	}

	return c.Status(code).JSON(result)
}

func ResponseErrorBadRequest(c *fiber.Ctx, message string) error {
	return responseError(c, fiber.StatusBadRequest, message)
}

func ResponseErrorNotFound(c *fiber.Ctx, message string) error {
	return responseError(c, fiber.StatusNotFound, message)
}

func ResponseErrorInternalServerError(c *fiber.Ctx, message string) error {
	return responseError(c, fiber.StatusInternalServerError, message)
}

func ResponseErrorUnauthorized(c *fiber.Ctx, message string) error {
	return responseError(c, fiber.StatusUnauthorized, message)
}

func ResponseErrorForbidden(c *fiber.Ctx, message string) error {
	return responseError(c, fiber.StatusForbidden, message)
}

func ResponseErrorConflict(c *fiber.Ctx, message string) error {
	return responseError(c, fiber.StatusConflict, message)
}
