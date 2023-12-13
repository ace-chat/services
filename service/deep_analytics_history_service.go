package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type DeepAnalyticsHistoryRequest struct{}

func (d *DeepAnalyticsHistoryRequest) GetHistory(user model.User) serializer.Response {
	analytics := make([]model.Analytics, 0)
	if err := cache.DB.Model(&model.Analytics{}).Where("user_id = ? AND type = ?", user.Id, 2).Find(&analytics).Error; err != nil {
		zap.L().Error("[DeepAnalytics] Get history failure", zap.Error(err))
		return serializer.DBError(err)
	}
	return serializer.Response{
		Code: 200,
		Data: analytics,
	}
}
