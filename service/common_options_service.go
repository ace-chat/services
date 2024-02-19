package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type OptionsRequest struct{}

func (r *OptionsRequest) GetOptions() serializer.Response {
	options := make([]model.BusinessChatBotOption, 0)
	if err := cache.DB.Model(&model.BusinessChatBotOption{}).Find(&options).Error; err != nil {
		zap.L().Error("[OptionsRequest] Find business bot option failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: 200,
		Data: options,
	}
}
