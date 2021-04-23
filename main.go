package main

import (
	"github.com/lnarasimhan83/go_rest/app"
	"github.com/lnarasimhan83/go_rest/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}