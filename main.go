package main

import (
	"./config"
	"./handler"
)

func main() {
	config := config.GetConfig()

	app := &handler.App{}
	app.Initialize(config)
	app.Run(":8090")
}
