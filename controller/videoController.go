package controller

import (
	"GitHub.com/LigeronAhill/GinLearning/entity"
	"GitHub.com/LigeronAhill/GinLearning/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context) entity.Video
	FindAll() []entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		return entity.Video{}
	}
	c.service.Save(video)
	return video

}
