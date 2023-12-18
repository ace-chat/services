package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
	"net/http"
)

type CommonVoicesRequest struct{}

func (c *CommonVoicesRequest) GetVoices(user model.User) serializer.Response {
	voices := make([]model.Voice, 0)
	if err := cache.DB.Model(&model.Voice{}).Where("user_id = ?", user.Id).Find(&voices).Error; err != nil {
		zap.L().Error("[Common] Get voices failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: voices,
	}
}
