package service

import (
	"ace/cache"
	"ace/model"
	"ace/request"
	"ace/serializer"
	"go.uber.org/zap"
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

	request.Client.Body = map[string]any{
		"brand_text": voice.Text,
	}
	body, err := request.Client.Post(model.Url["create_brand_voice"], false)
	if err != nil {
		zap.L().Error("[Common] Create brand voice failure", zap.Error(err))
		return serializer.GeneratorError(err)
	}
	voice.Content = string(body)

	if err := cache.DB.Model(&model.Voice{}).Create(&voice).Error; err != nil {
		zap.L().Error("[Common] Create voices failure", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: voice,
	}
}
