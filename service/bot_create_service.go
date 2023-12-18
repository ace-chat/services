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
	title := fmt.Sprintf("%v-%v", user.Id, time.Now().Unix())
	bot := model.ChatBot{
		UserId: user.Id,
		Title:  title,
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
