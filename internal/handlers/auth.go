package handlers

import (
	"net/http"

	"github.com/AronditFire/TODO-APP/internal/entities"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		SendError(c, http.StatusBadRequest, "invalid data while creating user")
	}

}
