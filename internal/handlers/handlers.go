package handlers

import (
	"github.com/AronditFire/TODO-APP/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHander(sv *service.Service) *Handler {
	return &Handler{service: sv}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signup)
		auth.POST("/sign-in" /*HANDLER*/)
		auth.POST("/refresh" /*HANDLER*/)
	}

	tasks := router.Group("/tasks" /*MIDDLEWARE*/)
	{
		tasks.GET("/" /*HANDLER*/)
		tasks.POST("/" /*HANDLER*/)
		tasks.PUT("/:id" /*HANDLER*/)
		tasks.DELETE("/:id" /*HANDLER*/)
	}
	admin := router.Group("/admin" /*MIDDLEWARE*/)
	{
		admin.POST("/" /*HANDLER*/)
	}

	return router
}
