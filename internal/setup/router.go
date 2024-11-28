package setup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/config/v2"
)

func SetupRouter(app *fiber.App) {
	// 挂载静态文件 到 /public
	app.Static("/public", "./resource/public")

	// 简单路由
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(config.String("AppName", "GoPress"))
	})

	// 参数路由
	app.Get("/test/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})

	// 可选参数路由
	app.Get("/test1/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("name: " + c.Params("name"))
		}
		return c.SendString("无参数")
	})

	// 通配符
	app.Get("/test2/*", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("*"))
	})
}
