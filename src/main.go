package main

import (
	"github.com/gin-gonic/gin"
	"go-media-fetcher/src/controllers"
	"go-media-fetcher/src/middleware"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	godotenv.Load(".env.local")
	godotenv.Load(".env")
}

func main() {
	if ginMode, exists := os.LookupEnv("GIN_MODE"); exists == true {
		gin.SetMode(ginMode)
	}

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.GET("/stream/get", controllers.StreamController{}.Get)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	_ = r.Run(":8090")
}