package main

import (
	"AuctionAPI/api/config"
	"AuctionAPI/api/handler"
)

func main() {
	config := config.GetConfig()

	app := &handler.App{}
	app.Initialize(config)
	app.Run(":8090")
}
