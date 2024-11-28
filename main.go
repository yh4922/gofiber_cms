package main

import (
	"fmt"
	"gopress/cmd"
	"log"

	"github.com/gookit/config/v2"
)

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
