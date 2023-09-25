package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Result struct {
	Data interface{} `json:"data"`
}

type PaginatedResult struct {
	Count   int         `json:"count"`
	HasMore bool        `json:"has_more"`
	Data    interface{} `json:"data"`
}

func Error(err error, code int, message string) *fiber.Error {
	log.Print(err)
	return fiber.NewError(code, message)
}
