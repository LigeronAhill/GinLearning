package main

import (
	"GitHub.com/LigeronAhill/GinLearning/controller"
	"GitHub.com/LigeronAhill/GinLearning/middlewares"
	"GitHub.com/LigeronAhill/GinLearning/service"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func setupLogOutput() {
	if file, err := os.Create("gin.log"); err == nil {
		gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	}
}

func main() {

	setupLogOutput()

	videoService := service.New()
	videoController, err := controller.New(videoService)
	if err != nil {
		log.Fatal(err)
	}

	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger())

	apiRoutes := server.Group("/api", middlewares.BasicAuth())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			if err := videoController.Save(ctx); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := server.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
