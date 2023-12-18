package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type BotCreate struct{}

func (b *BotCreate) Create(user model.User) serializer.Response {
	chatId := fmt.Sprintf("%v-%v", user.Id, time.Now().Unix())
	title := fmt.Sprintf("chat bot for %v", user.Id)
	bot := model.ChatBot{
		UserId: user.Id,
		Title:  title,
		ChatId: chatId,
	}
	if err := cache.DB.Model(&model.ChatBot{}).Create(&bot).Error; err != nil {
		zap.L().Error("[ChatBot] Create chat bot failed", zap.Error(err))
		return serializer.DBError(err)
	}
	return serializer.Response{
		Code: 200,
		Data: bot,
	}
}
