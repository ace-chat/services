package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type IntroHistoryRequest struct{}

func (t *IntroHistoryRequest) GetHistory(user model.User) serializer.Response {
	histories := make([]model.BlogContent, 0)
	if err := cache.DB.Model(&model.BlogContent{}).Where("user_id = ? AND type = ?", user.Id, 1).Find(&histories).Error; err != nil {
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: histories,
	}
}
