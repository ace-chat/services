package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
	"net/http"
)

type LanguageRequest struct{}

func (l *LanguageRequest) GetLanguages() serializer.Response {
	languages := make([]model.Language, 0)
	if err := cache.DB.Model(&model.Language{}).Find(&languages).Error; err != nil {
		zap.L().Error("[Common] Get language failure", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: languages,
	}
}
