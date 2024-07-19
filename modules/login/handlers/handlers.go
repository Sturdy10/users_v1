package handlers

import (
	"net/http"
	"users_v1/modules/login/models"
	"users_v1/modules/login/services"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	LoginH(c *gin.Context)
}

type handler struct {
	s services.IService
}

func NewHandler(s services.IService) IHandler {
	return &handler{s: s}
}

func (h *handler) LoginH(c *gin.Context) {
	var login models.LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return 
	}
		err := h.s.LoginS(login) 
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()} )
			return
		}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Login successfully"})
}