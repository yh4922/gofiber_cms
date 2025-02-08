package ctx

import (
	"github.com/gofiber/fiber/v2"
)

func ThemeInfo(c *fiber.Ctx) error {
	// Create data map for template rendering
	data := fiber.Map{
		"name": "GoPress",
	}

	// Render index.html template with data
	return c.Render("resource/template/info.html", data)
}