package main

import (
	"echobox/config"
	"echobox/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	routes.RegisterRoutes(r)
	err := r.Run(fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		slog.Error("Error launching server.")
		os.Exit(1)
	}
}
