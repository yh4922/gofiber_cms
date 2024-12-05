package ctx

import (
	"github.com/gofiber/fiber/v2"
)

func ManageUserLoginInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "success"})
}
