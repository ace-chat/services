package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type TypesRequest struct{}

func (g *TypesRequest) GetTypes() serializer.Response {
	types := make([]model.Type, 0)
	if err := cache.DB.Model(&model.Type{}).Find(&types).Error; err != nil {
		zap.L().Error("[Common] Get types failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: 200,
		Data: types,
	}
}
