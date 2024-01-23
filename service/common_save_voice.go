package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"

	"go.uber.org/zap"
)

type CommonSaveVoiceRequest struct {
	Name    string `form:"name" json:"name" binding:"required"`
	Text    string `form:"text" json:"text" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

func (c *CommonSaveVoiceRequest) SaveVoice(user model.User) serializer.Response {
	voice := model.Voice{
		UserId:  user.Id,
		Name:    c.Name,
		Text:    c.Text,
		Content: c.Content,
	}

	if err := cache.DB.Model(&model.Voice{}).Create(&voice).Error; err != nil {
		zap.L().Error("[Common] Save brand voice failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: voice,
	}
}
