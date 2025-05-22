package route

import (
	"github.com/gin-gonic/gin"
	"ragnarok/internal/route/api"
)

func RegisterRoutes(r *gin.Engine) {
	// Healthcheck
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Grouped APIs
	api.RegisterBrimoExpeditionRoutes(r)
}
