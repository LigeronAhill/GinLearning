package controller

import (
	"GitHub.com/LigeronAhill/GinLearning/entity"
	"GitHub.com/LigeronAhill/GinLearning/service"
	"GitHub.com/LigeronAhill/GinLearning/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) (VideoController, error) {
	validate = validator.New()
	err := validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	if err != nil {
		return nil, err
	}
	return &controller{
		service: service,
	}, nil
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil

}
