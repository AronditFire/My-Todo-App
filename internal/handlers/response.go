package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func SendError(c *gin.Context, status int, message string) {
	log.Fatal(message)
	c.AbortWithStatusJSON(status, ErrorMessage{message})
}
