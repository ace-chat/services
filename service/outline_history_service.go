package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
	"net/http"
)

type OutlineHistoryRequest struct{}

func (t *OutlineHistoryRequest) GetHistory(user model.User) serializer.Response {
	histories := make([]model.BlogContent, 0)
	if err := cache.DB.Model(&model.BlogContent{}).Where("user_id = ? AND type = ?", user.Id, 2).Find(&histories).Error; err != nil {
		zap.L().Error("[Outline] Get blog history failure", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: histories,
	}
}
