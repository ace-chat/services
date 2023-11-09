package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type EngineHistoryRequest struct{}

func (e *EngineHistoryRequest) GetHistory(user model.User) serializer.Response {
	histories := make([]model.EngineContent, 0)
	if err := cache.DB.Model(&model.EngineContent{}).Where("user_id = ?", user.Id).Find(&histories).Error; err != nil {
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: histories,
	}
}
