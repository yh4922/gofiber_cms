package cmd

import (
	"gopress/internal/setup"

	"github.com/gookit/config/v2"

	"github.com/gofiber/fiber/v2"
)

// 错误处理
func errorHandler(ctx *fiber.Ctx, err error) error {
	// 状态码 默认500
	code := fiber.StatusInternalServerError

	// 获取自定义状态码
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Send custom error page
	return ctx.Status(code).JSON(fiber.Map{
		"code": code,
		"msg":  err.Error(),
	})
}

var (
	App *fiber.App
)

// 初始化
func init() {
	setup.SetupConfig()

	// 初始化 Fiber 应用
	App = fiber.New(fiber.Config{
		AppName:      config.String("AppName", "GoPress"),
		ErrorHandler: errorHandler,
	})

	// 配置路由
	setup.SetupRouter(App)
}
