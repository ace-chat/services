package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type BusinessDeleteRequest struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

func (r *BusinessDeleteRequest) Delete(user model.User) serializer.Response {
	var count int64
	if err := cache.DB.Model(&model.BusinessChatBot{}).Where("user_id = ?", user.Id).Count(&count).Error; err != nil {
		zap.L().Error("[BusinessDeleteRequest] Get business chat bot count failed", zap.Error(err))
		return serializer.DBError(err)
	}

	if count == 0 {
		return serializer.IllegalError()
	}

	var chat model.BusinessChatBot
	if err := cache.DB.Model(&model.BusinessChatBot{}).Where("user_id = ?", user.Id).Last(&chat).Error; err != nil {
		zap.L().Error("[BusinessDeleteRequest] Get business chat bot failed", zap.Error(err))
		return serializer.DBError(err)
	}

	// remove business chat bot
	if err := cache.DB.Model(&model.BusinessChatBot{}).Delete(&chat).Error; err != nil {
		zap.L().Error("[BusinessDeleteRequest] Delete business chat bot failed", zap.Error(err))
		return serializer.DBError(err)
	}

	// remove business chat bot files
	fs := model.BusinessChatBotFile{
		BusinessBotId: chat.Id,
	}
	if err := cache.DB.Model(&model.BusinessChatBotFile{}).Delete(&fs).Error; err != nil {
		zap.L().Error("[BusinessDeleteRequest] Delete business chat bot files failed", zap.Error(err))
		return serializer.DBError(err)
	}

	// remove business chat bot links
	links := model.BusinessChatBotLink{
		BusinessBotId: chat.Id,
	}
	if err := cache.DB.Model(&model.BusinessChatBotLink{}).Delete(&links).Error; err != nil {
		zap.L().Error("[BusinessDeleteRequest] Delete business chat bot links failed", zap.Error(err))
		return serializer.DBError(err)
	}

	// remove business chat bot question and answer
	qa := model.BusinessChatBotQA{
		BusinessBotId: chat.Id,
	}
	if err := cache.DB.Model(&model.BusinessChatBotQA{}).Delete(&qa).Error; err != nil {
		zap.L().Error("[BusinessDeleteRequest] Delete business chat bot question and answer failed", zap.Error(err))
		return serializer.DBError(err)
	}

	// remove business chat bot sales and pitches
	sp := model.BusinessChatBotSalesPitch{
		BusinessBotId: chat.Id,
	}
	if err := cache.DB.Model(&model.BusinessChatBotSalesPitch{}).Delete(&sp).Error; err != nil {
		zap.L().Error("[BusinessDeleteRequest] Delete business chat bot sales and pitches failed", zap.Error(err))
		return serializer.DBError(err)
	}

	// remove business chat bot platform
	p := model.BusinessChatBotPlatform{
		BusinessBotId: chat.Id,
	}
	if err := cache.DB.Model(&model.BusinessChatBotPlatform{}).Delete(&p).Error; err != nil {
		zap.L().Error("[BusinessDeleteRequest] Delete business chat bot platform failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
	}
}
