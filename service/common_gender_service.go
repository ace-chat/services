package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type GenderRequest struct{}

func (g *GenderRequest) GetGender() serializer.Response {
	genders := make([]model.Gender, 0)
	if err := cache.DB.Model(&model.Gender{}).Find(&genders).Error; err != nil {
		zap.L().Error("[Common] Get genders failure", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: 200,
		Data: genders,
	}
}
