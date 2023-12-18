package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type ServicesRequest struct{}

func (s *ServicesRequest) GetServices() serializer.Response {
	services := make([]model.Service, 0)
	if err := cache.DB.Model(&model.Service{}).Find(&services).Error; err != nil {
		zap.L().Error("[Common] Get service failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: 200,
		Data: services,
	}
}
