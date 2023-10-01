package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Result struct {
	Data interface{} `json:"data"`
}

type PaginatedResult[T any] struct {
	Count   int  `json:"count"`
	HasMore bool `json:"has_more"`
	Data    []T  `json:"data"`
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

func Error(err error, code int, message string) *fiber.Error {
	log.Print(err)
	return fiber.NewError(code, message)
}
