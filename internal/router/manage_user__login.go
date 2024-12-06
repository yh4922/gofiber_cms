package router

import (
	ctx "gopress/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerItem := RouterItem{
		Path:   "/api/manage/user/login",
		Method: "GET",
		Func:   ctx.ManageUserLogin,
	}

	routerItem.Middle = []func(c *fiber.Ctx) error{}

	RouterList = append(RouterList, routerItem)
}
