package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type AudienceHistoryRequest struct{}

func (t *AudienceHistoryRequest) GetHistory(user model.User) serializer.Response {
	histories := make([]model.OptimizedContent, 0)
	if err := cache.DB.Model(&model.OptimizedContent{}).Where("user_id = ? AND type = ?", user.Id, 5).Find(&histories).Error; err != nil {
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: histories,
	}
}
