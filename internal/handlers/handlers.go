package handlers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", signup)
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
