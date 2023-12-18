package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type BotChatList struct{}

func (b *BotChatList) GetChatList(user model.User) serializer.Response {
	bots := make([]model.ChatBot, 0)
	if err := cache.DB.Model(&model.ChatBot{}).Where("user_id = ?", user.Id).Find(&bots).Error; err != nil {
		zap.L().Error("[ChatBot] Get chat bot list failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: 200,
		Data: bots,
	}
}
