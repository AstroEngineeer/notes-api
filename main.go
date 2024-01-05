package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vigneshganesan008/notes-api/api"
	"github.com/vigneshganesan008/notes-api/dao"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dao.InitDb()
	defer dao.CloseDB()

	router := gin.Default()

	authGroup := router.Group("/auth")
	authGroup.POST("/signup", api.Signup)
	authGroup.POST("/login", api.Login)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
