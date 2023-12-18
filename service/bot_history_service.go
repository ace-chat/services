package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BotHistory struct {
	Id *uint `form:"id" json:"id" binding:"required"`
}

func (b *BotHistory) GetHistory(user model.User) serializer.Response {
	var bot model.ChatBot
	if err := cache.DB.Model(&model.ChatBot{}).Where("id = ? AND user_id = ?", *b.Id, user.Id).First(&bot).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		zap.L().Error("[ChatBot] Get chat bot history failed", zap.Error(err))
		return serializer.DBError(err)
	}

	histories := make([]any, 0)
	//ctx := context.Background()
	//cursor, err := cache.Mongo.Collection(bot.Title).Find(ctx, nil, nil)
	//if err != nil {
	//	return serializer.Response{}
	//}
	return serializer.Response{
		Code: 200,
		Data: histories,
	}
}