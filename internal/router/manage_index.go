package router

import (
	"github.com/gofiber/fiber/v2"
	"gopress/internal/controller"
)


func init() {
	routerItem := RouterItem{
		Path:   "/",
		Method: "GET",
		Middle: []func(c *fiber.Ctx) error{},
		Func:   ctx.Index,
	}

	routerItem.Middle = []func(c *fiber.Ctx) error{}

	RouterList = append(RouterList, routerItem)
}
