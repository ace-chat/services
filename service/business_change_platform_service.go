package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type ChangeBusinessPlatformRequest struct {
	Id     *uint `form:"id" json:"id" binding:"required"`
	Status *bool `form:"status" json:"status" binding:"required"`
}

func (r *ChangeBusinessPlatformRequest) Change(user model.User) serializer.Response {
	var businessChatBotPlatform model.BusinessChatBotPlatform
	if err := cache.DB.Model(&model.BusinessChatBotPlatform{}).Where("id = ?", *r.Id).Last(&businessChatBotPlatform).Error; err != nil {
		zap.L().Error("[ChangeBusinessPlatformRequest] Get business chat bot platform failed", zap.Error(err))
		return serializer.DBError(err)
	}

	var platform model.Platform
	if err := cache.DB.Model(&model.Platform{}).Where("id = ?", businessChatBotPlatform.Platform).First(&platform).Error; err != nil {
		zap.L().Error("[ChangeBusinessPlatformRequest] Get platform failed", zap.Error(err))
		return serializer.DBError(err)
	}

	if !platform.Status {
		return serializer.IllegalError()
	}

	var count int64
	if err := cache.DB.Model(&model.BusinessChatBot{}).Where("user_id = ?", user.Id).Count(&count).Error; err != nil {
		zap.L().Error("[ChangeBusinessPlatformRequest] Get business chat bot count failed", zap.Error(err))
		return serializer.DBError(err)
	}

	if count == 0 {
		return serializer.IllegalError()
	}

	if err := cache.DB.Model(&model.BusinessChatBotPlatform{}).Where("id = ?", businessChatBotPlatform.Id).Update("status", *r.Status).Error; err != nil {
		zap.L().Error("[ChangeBusinessPlatformRequest] Update business chat bot platform status failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
	}
}
