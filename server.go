package main

import (
	"GitHub.com/LigeronAhill/GinLearning/controller"
	"GitHub.com/LigeronAhill/GinLearning/middlewares"
	"GitHub.com/LigeronAhill/GinLearning/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
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

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")
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

	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
