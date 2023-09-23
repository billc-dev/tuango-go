package utils

import "github.com/gofiber/fiber/v2"

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

func NewError(ctx *fiber.Ctx, status int, err error) {
	ctx.
		Status(status).
		JSON(HTTPError{
			Code:    status,
			Message: err.Error(),
		})
}
