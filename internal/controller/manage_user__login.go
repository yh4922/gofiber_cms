package ctx

import (
	"github.com/gofiber/fiber/v2"
)

func ManageUserLogin(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "success"})
}
