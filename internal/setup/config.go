package setup

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json5"
)

func SetupConfig() {
	config.AddDriver(json5.Driver)

	err := config.LoadFiles("./configs/config.json")
	if err != nil {
		panic(err)
	}
}
