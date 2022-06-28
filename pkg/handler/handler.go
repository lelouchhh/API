package handler

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"AGZ/pkg/service"

	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.GET("/sign-up", h.signUp)
		auth.GET("/sign-in", h.signIn)
	}
	token := router.Group("/token")
	{
		token.GET("/actionAccess", h.ActionUserAccess)
		token.GET("/actionRefresh", h.ActionUserRefresh)
	}
	procedures :=  router.Group("/procedures")
	{
		procedures.GET("create_news", h.CreateNews)
		procedures.GET("get_entity", h.GetEntity)
	}
	return router
}