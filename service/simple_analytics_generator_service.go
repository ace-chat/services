package service

import (
	"ace/cache"
	"ace/model"
	"ace/request"
	"ace/serializer"
	"fmt"
	"go.uber.org/zap"
)

type SimpleAnalytics struct {
	Filename *string `form:"filename" json:"filename" binding:"required"`
}

func (s *SimpleAnalytics) Generator(user model.User) serializer.Response {
	analytics := model.Analytics{
		UserId: user.Id,
		Title:  fmt.Sprintf("Simple analytics for %v", *s.Filename),
		Type:   1,
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

	simpleAnalytics := serializer.SimpleAnalytics{
		Id:        analytics.Id,
		Title:     analytics.Title,
		Content:   string(body),
		CreatedAt: analytics.CreatedAt,
	}

	return simpleAnalytics.Bind()
}