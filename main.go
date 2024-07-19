package main

import (
	"log"
	"users_v1/modules/servers/login"
	"users_v1/modules/servers/password"
	"users_v1/modules/servers/register"
	"users_v1/pkg/databases/postgresql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := postgresql.Postgresql()
	defer db.Close()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"}
	router.Use(cors.New(config))

	login.Login(router, db)
	password.Password(router, db)
	register.Register(router, db)

	err := router.Run(":8888")
	if err != nil {
		log.Fatal(err.Error())
	}
}
