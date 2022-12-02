package main

import (
	"todo-list/config"
	"todo-list/routes"
)

func main() {
	config.Init()
	r := routes.NewRouter()
	_ = r.Run(config.HttpPort)
}
