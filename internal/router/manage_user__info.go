package router

import (
	ctx "gopress/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerItem := RouterItem{
		Path:   "/info/:id",
		Method: "GET",
		Func:   ctx.ThemeInfo,
	}

	routerItem.Middle = []func(c *fiber.Ctx) error{}

	RouterList = append(RouterList, routerItem)
}