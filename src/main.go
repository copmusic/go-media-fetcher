package main

import (
	"github.com/gin-gonic/gin"
	"go-media-fetcher/src/controllers"
	"go-media-fetcher/src/middleware"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env.local"); err != nil {
		panic("No .env.local file found")
	}
}

func main() {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.GET("/stream/get", controllers.StreamController{}.Get)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	_ = r.Run(":8080")
}