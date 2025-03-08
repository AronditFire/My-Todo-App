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

func (h *Handler) signin(c *gin.Context) {
	var req entities.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendError(c, http.StatusBadRequest, "invalid data in login")
		return
	}
	user, err := h.service.Authorization.CheckUser(req)
	if err != nil {
		SendError(c, http.StatusBadRequest, "incorrent login or password")
	}

	accessToken, err := h.service.Authorization.GenerateAccessToken(user)
	if err != nil {
		SendError(c, http.StatusInternalServerError, err.Error())
	}

	refreshToken, err := h.service.Authorization.GenerateRefreshToken(user)
	if err != nil {
		SendError(c, http.StatusInternalServerError, err.Error())
	}

	c.SetCookie("accessToken", accessToken, 15*60, "/", "", false, true)         // TODO: change false to true if https
	c.SetCookie("refreshToken", refreshToken, 30*24*60*60, "/", "", false, true) // TODO: change false to true if https

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in",
	})
}

func (h *Handler) RefreshAccessToken(c *gin.Context) {
	refreshTokenString, err := c.Cookie("refreshToken") // maybe postform?
	if err != nil {
		SendError(c, http.StatusUnauthorized, "refresh token not found")
		return
	}

	newAccessToken, err := h.service.Authorization.RefreshAccessToken(refreshTokenString)
	if err != nil {
		SendError(c, http.StatusUnauthorized, "refresh token is invalid")
	}

	c.SetCookie("accessToken", newAccessToken, 15*60, "/", "", false, true)
}
