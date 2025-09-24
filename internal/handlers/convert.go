package handlers

import (
	"github.com/gofiber/fiber/v2"
	"mastery-backend-go/internal/utils"
	"mastery-backend-go/internal/settings"
)

func ConvertHandler(c *fiber.Ctx) error {
	url := c.Query("url")
	if url == "" {
		return c.JSON(fiber.Map{"success": false, "error": "URL required"})
	}

	cfg := settings.LoadSettings()
	data, err := utils.ExtractInfo(url, "bestaudio/best", cfg.YtCookies)
	if err != nil {
		return c.JSON(data)
	}
	return c.JSON(data)
}
