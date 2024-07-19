package register

import (
	"database/sql"
	"users_v1/modules/register/handlers"
	"users_v1/modules/register/repositories"
	"users_v1/modules/register/services"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, db *sql.DB) {

	r := repositories.NewRepositorie(db)
	s := services.NewService(r)
	h := handlers.NewHandler(s)

	router.POST("/api/register", h.RegisterMemberHandler)
	router.GET("/api/getAllMember", h.GetallMembersHandler)
}
