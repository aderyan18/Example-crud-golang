package main

import (
	"belajar-go/config"
	"belajar-go/controllers"
	"belajar-go/model"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init () {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDatabase()
	config.DB.AutoMigrate(&model.Todo{})
}

	func main() {
		router := gin.Default()

		router.GET("/todo", controllers.Index)
		router.POST("/todo", controllers.Create)
		router.GET("/todo/:id", controllers.Show)
		router.PUT("/todo/:id", controllers.Update)
		router.DELETE("/todo/:id", controllers.Delete)

		router.Run(":" + os.Getenv("APP_PORT"))
	}