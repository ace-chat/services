package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type ToneHistoryRequest struct{}

func (t *ToneHistoryRequest) GetHistory(user model.User) serializer.Response {
	histories := make([]model.OptimizedContent, 0)
	if err := cache.DB.Model(&model.OptimizedContent{}).Where("user_id = ? AND type = ?", user.Id, 1).Find(&histories).Error; err != nil {
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: histories,
	}
}
