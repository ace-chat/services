package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type SimpleAnalyticsHistoryRequest struct{}

func (s *SimpleAnalyticsHistoryRequest) GetHistory(user model.User) serializer.Response {
	analytics := make([]model.Analytics, 0)
	if err := cache.DB.Model(&model.Analytics{}).Where("user_id = ? AND type = ?", user.Id, 1).Find(&analytics).Error; err != nil {
		zap.L().Error("[Analytics] Get history failed", zap.Error(err))
		return serializer.DBError(err)
	}
	return serializer.Response{
		Code: 200,
		Data: analytics,
	}
}
