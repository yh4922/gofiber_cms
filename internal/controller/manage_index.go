package ctx

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	// Create data map for template rendering
	data := fiber.Map{
		"name": "GoPress",
	}

	// Render index.html template with data
	return c.Render("resource/template/index.html", data)
}