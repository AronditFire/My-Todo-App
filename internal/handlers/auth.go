package handlers

import (
	"net/http"

	"github.com/AronditFire/TODO-APP/internal/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signup(c *gin.Context) {
	var user entities.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		SendError(c, http.StatusBadRequest, "invalid data while creating user")
		return
	}

	id, err := h.service.Authorization.CreateUser(user)
	if err != nil {
		SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})

}
