package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func main() {
	var app *fiber.App = fiber.New(fiber.Config{
		CaseSensitive: true,
		ErrorHandler: func(c fiber.Ctx, err error) error {
			fmt.Println("fiber error")
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		},
	})
}