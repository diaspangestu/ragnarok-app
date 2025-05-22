package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ragnarok/config"
	"ragnarok/internal/route"
)

func StartServer() error {
	r := gin.Default()

	// Register route
	route.RegisterRoutes(r)

	addr := fmt.Sprintf(":%s", config.Cfg.Port)
	fmt.Printf("🚀 Server running at http://localhost%s\n", addr)

	return r.Run(addr) // start the Gin server
}
