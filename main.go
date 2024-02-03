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

	authGroup := router.Group("/api/auth")
	authGroup.POST("/signup", api.Signup)
	authGroup.POST("/login", api.Login)

	// noteGroup := router.Group("/api")
	// noteGroup.GET("/notes", api.ListNotes)
	// noteGroup.GET("/notes/:id", api.GetNote)
	// noteGroup.POST("/notes", api.CreateNote)
	// noteGroup.PUT("/notes/:id", api.UpdateNote)
	// noteGroup.DELETE("/notes/:id", api.DeleteNote)
	// noteGroup.POST("/notes/:id/share", api.ShareNote)
	// noteGroup.GET("/notes/search", api.SeachNotes)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
