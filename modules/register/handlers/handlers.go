package handlers

import (
	"net/http"
	"users_v1/modules/register/models"
	"users_v1/modules/register/services"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	RegisterMemberHandler(c *gin.Context)
	GetallMembersHandler(c *gin.Context)
}

type handler struct {
	s services.IService
}

func NewHandler(s services.IService) IHandler {
	return &handler{s: s}
}

func (h *handler) RegisterMemberHandler(c *gin.Context) {
	var addMember models.RegisterMember
	if err := c.ShouldBindJSON(&addMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	err := h.s.RegisterMemberService(addMember)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "OK", "message": "Register Successfully"})
}

func (h *handler) GetallMembersHandler(c *gin.Context) {
	members, err := h.s.GetallMemberService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "DataMembers": members})
}
