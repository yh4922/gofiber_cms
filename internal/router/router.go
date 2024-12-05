package router

import (
	"github.com/gofiber/fiber/v2"
)

type RouterItem struct {
	Path   string                     // 路径 必填的
	Method string                     // 方法 必填的
	Middle []func(c *fiber.Ctx) error // 中间件函数列表
	Func   func(c *fiber.Ctx) error   // 函数 必填的
}

var RouterList = []RouterItem{}
