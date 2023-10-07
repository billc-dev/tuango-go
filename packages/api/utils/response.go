package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Result[T any] struct {
	Data T `json:"data"`
}

type InfinitePaginatedResult[T any] struct {
	HasMore bool `json:"has_more"`
	Data    []T  `json:"data"`
}

type PaginatedResult[T any] struct {
	Count   int  `json:"count"`
	HasMore bool `json:"has_more"`
	Data    []T  `json:"data"`
}

type HTTPError string

func Error(err error, code int, message string) *fiber.Error {
	log.Println(err)
	return fiber.NewError(code, message)
}
