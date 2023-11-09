package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type FreestyleHistoryRequest struct{}

func (t *FreestyleHistoryRequest) GetHistory(user model.User) serializer.Response {
	histories := make([]model.EmailContent, 0)
	if err := cache.DB.Model(&model.EmailContent{}).Where("user_id = ? AND type = ?", user.Id, 1).Find(&histories).Error; err != nil {
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: histories,
	}
}
