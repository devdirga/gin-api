package main

import (
	"go/gin-api/controllers"
	"go/gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "pong",
		})
	})
	r.GET("books", controllers.Finds)
	r.GET("book/:id", controllers.Find)
	r.POST("book", controllers.Create)
	r.PUT("book/:id", controllers.Update)
	r.DELETE("book/:id", controllers.Delete)
	r.POST("upload", controllers.SaveFileHandler)
	r.POST("email", controllers.SendMail)
	r.Run()
}
