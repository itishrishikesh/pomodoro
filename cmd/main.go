package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"main.go/internal/config"
	"main.go/internal/controller"
)

func main() {
	config, err := config.Load()
	if err != nil {
		panic("Failed to load config")
	}
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.Run(fmt.Sprintf(":%d", config.Port))
}
