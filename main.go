package main

import (
	"fmt"
	"gopress/cmd"
	"log"

	"github.com/gookit/config/v2"
)

// @title goPress测试项目
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	// 启动服务
	log.Fatal(cmd.App.Listen(
		fmt.Sprintf(
			"%s:%d",
			config.String("Server.Host", "0.0.0.0"),
			config.Int("Server.Port", 3000),
		),
	))
}
