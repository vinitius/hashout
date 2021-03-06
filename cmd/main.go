package main

import (
	"viniti.us/hashout/config/log"

	"viniti.us/hashout/cmd/app"
)

func init() {
	app.Bootstrap("./")
}

func main() {
	err := app.Run() // this can be wrapped in a different routine
	if err != nil {
		log.Logger.Fatalw("Error starting http server", "error", err)
	}
}
