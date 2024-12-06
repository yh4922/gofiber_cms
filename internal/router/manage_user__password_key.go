package router

import (
	ctx "gopress/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerItem := RouterItem{
		Path:   "/api/manage/user/password-key",
		Method: "GET",
		Func:   ctx.ManageUserPasswordKey,
	}

	routerItem.Middle = []func(c *fiber.Ctx) error{}

	RouterList = append(RouterList, routerItem)
}
