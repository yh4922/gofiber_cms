package ctx

import (
	"github.com/gofiber/fiber/v2"
)

func ManageUserPasswordKey(c *fiber.Ctx) error {
	rsaKey, publicKey, err := GetRsaKey()
	if err != nil {
		return c.JSON(CtxError(500, err.Error()))
	}

	return c.JSON(CtxSuccess(fiber.Map{
		"key":   rsaKey,
		"value": publicKey,
	}))
}
