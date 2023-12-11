package service

import (
	"ace/cache"
	"ace/model"
	"ace/request"
	"ace/serializer"
	"go.uber.org/zap"
)

type SimpleAnalytics struct {
	Filename *string `form:"filename" json:"filename" binding:"required"`
}

func (s *SimpleAnalytics) Generator() serializer.Response {
	analytics := model.Analytics{
		Type: 1,
	}

	body, err := request.Client.Post(*s.Filename, true)
	if err != nil {
		zap.L().Error("[Analytics] Create simple analytics failure", zap.Error(err))
		return serializer.GeneratorError(err)
	}
	analytics.Content = string(body)
	if err := cache.DB.Model(&model.Analytics{}).Create(&analytics).Error; err != nil {
		zap.L().Error("[Analytics] Create simple analytics record failure", zap.Error(err))
		return serializer.DBError(err)
	}

	var simpleAnalytics serializer.SimpleAnalytics

	return simpleAnalytics.Bind(string(body))
}
