package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
	"net/http"
)

type AudienceHistoryRequest struct{}

func (t *AudienceHistoryRequest) GetHistory(user model.User) serializer.Response {
	histories := make([]model.OptimizedContent, 0)
	if err := cache.DB.Model(&model.OptimizedContent{}).Where("user_id = ? AND type = ?", user.Id, 5).Find(&histories).Error; err != nil {
		zap.L().Error("[Audience] Get optimized history failure", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: histories,
	}
}
