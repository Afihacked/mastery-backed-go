package main

import (
	"github.com/gofiber/fiber/v2"
	"mastery-backend-go/internal/handlers"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "MasterY Backend (Go) is running 🚀"})
	})

	api := app.Group("/api")
	api.Get("/extract", handlers.ExtractHandler)
	api.Get("/convert", handlers.ConvertHandler)

	app.Listen(":10000")
}
