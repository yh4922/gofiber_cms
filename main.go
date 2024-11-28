package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// 静态文件
	app.Static("/public", "./resource/public")

	// 简单路由
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
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

	log.Fatal(app.Listen(":3000"))
}
