package main

import (
	"github.com/gofiber/fiber/v2"
	"mastery-backend-go/internal/handlers"
)

func main() {
	app := fiber.New()
// Debug endpoint buat cek yt-dlp
    app.Get("/debug/yt", func(c *fiber.Ctx) error {
        path, err := exec.LookPath("yt-dlp")
        if err != nil {
            return c.JSON(fiber.Map{"error": "yt-dlp not found"})
        }
        out, _ := exec.Command(path, "--version").Output()
        return c.JSON(fiber.Map{
            "path":    path,
            "version": string(out),
        })
    })
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "MasterY Backend (Go) is running 🚀"})
	})

	api := app.Group("/api")
	api.Get("/extract", handlers.ExtractHandler)
	api.Get("/convert", handlers.ConvertHandler)

	app.Listen(":10000")
}
