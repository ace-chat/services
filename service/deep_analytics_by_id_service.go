package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DeepAnalyticsByIdRequest struct {
	Id *uint `form:"id" json:"id" binding:"required"`
}

func (s *DeepAnalyticsByIdRequest) GetAnalytics(user model.User) serializer.Response {
	var analytics model.Analytics
	if err := cache.DB.Model(&model.Analytics{}).Where("user_id = ? AND id = ? AND type = ?", user.Id, *s.Id, 2).First(&analytics).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		zap.L().Error("[DeepAnalytics] Get analytics content failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: 200,
		Data: analytics,
	}
}
