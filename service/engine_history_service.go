package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
	"net/http"
)

type EngineHistoryRequest struct{}

func (e *EngineHistoryRequest) GetHistory(user model.User) serializer.Response {
	histories := make([]model.EngineContent, 0)
	if err := cache.DB.Model(&model.EngineContent{}).Where("user_id = ?", user.Id).Find(&histories).Error; err != nil {
		zap.L().Error("[Engine] Get search engine ads history failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: histories,
	}
}
