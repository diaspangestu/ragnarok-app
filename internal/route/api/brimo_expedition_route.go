package api

import (
	"github.com/gin-gonic/gin"
	"ragnarok/internal/controller"
	"ragnarok/internal/repository"
	"ragnarok/internal/service"
)

func RegisterBrimoExpeditionRoutes(router *gin.Engine) {
	// Setup Dependencies
	repo := repositories.NewBrimoExpeditionRepo()
	service := service.NewBrimoExpeditionService(repo)
	controller := controller.NewBrimoExpeditionController(service)

	brimoExpedition := router.Group("/")
	{
		brimoExpedition.GET("/brimo-expeditions", controller.GetListBrimoExpedition)
	}
}
