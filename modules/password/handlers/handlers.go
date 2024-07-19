package handlers

import (
	"net/http"
	"users_v1/modules/password/models"
	"users_v1/modules/password/services"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	InitPasswordHandler(c *gin.Context)
	ChangePasswordHandler(c *gin.Context)
}

type handler struct {
	s services.IService
}

func NewHandler(s services.IService) IHandler {
	return &handler{s: s}
}

func (h *handler) InitPasswordHandler(c *gin.Context) {
	var initPassword models.InitPassword
	if err := c.ShouldBindJSON(&initPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	err := h.s.InitPasswordService(initPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Password initialized successfully"})
}
func (h *handler) ChangePasswordHandler(c *gin.Context) {
	var changePassword models.ChangePassword
	if err := c.ShouldBindJSON(&changePassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	err := h.s.ChangePasswordService(changePassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Password Changed successfully"})
}
