package router

import (
	ctx "gopress/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerItem := RouterItem{
		Path:   "/api/manage/user/login_info",
		Method: "GET",
		Func:   ctx.ManageUserLoginInfo,
	}

	routerItem.Middle = []func(c *fiber.Ctx) error{}

	RouterList = append(RouterList, routerItem)
}
