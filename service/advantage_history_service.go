package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type AdvantageHistoryRequest struct{}

func (t *AdvantageHistoryRequest) GetHistory(user model.User) serializer.Response {
	histories := make([]model.EmailContent, 0)
	if err := cache.DB.Model(&model.EmailContent{}).Where("user_id = ? AND type = ?", user.Id, 4).Find(&histories).Error; err != nil {
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: histories,
	}
}
