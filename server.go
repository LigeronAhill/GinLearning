package main

import (
	"GitHub.com/LigeronAhill/GinLearning/controller"
	"GitHub.com/LigeronAhill/GinLearning/service"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
