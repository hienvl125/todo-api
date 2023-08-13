// cmd/api/main.go
package main

import (
	"log"

	"github.com/hienvl125/todo-api/internal/registry"
	"github.com/hienvl125/todo-api/internal/util/config"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	ginServer, err := registry.NewGinServer(conf)
	if err != nil {
		log.Fatalln(err)
	}

	if err := ginServer.Run(conf.GinPort); err != nil {
		log.Fatalln(err)
	}
}
