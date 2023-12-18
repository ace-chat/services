package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
	"net/http"
)

type MediaHistoryRequest struct{}

func (m *MediaHistoryRequest) GetHistory(user model.User) serializer.Response {
	histories := make([]model.MediaContent, 0)
	if err := cache.DB.Model(&model.MediaContent{}).Where("user_id = ?", user.Id).Find(&histories).Error; err != nil {
		zap.L().Error("[Media] Get social media history failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: histories,
	}
}
