package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type CommonCreateVoiceRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
	Text string `form:"text" json:"text" binding:"required"`
}

func (c *CommonCreateVoiceRequest) CreateVoice(user model.User) serializer.Response {
	voice := model.Voice{
		UserId: user.Id,
		Name:   c.Name,
		Text:   c.Text,
	}

	if err := cache.DB.Model(&model.Voice{}).Create(&voice).Error; err != nil {
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: voice,
	}
}
