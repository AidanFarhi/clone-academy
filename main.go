package main

import (
	"clone-academy/app"
	"clone-academy/app/config"
)

func main() {
	config.InitDatabase()
	app.Run()
}
