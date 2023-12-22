package service

import (
	"ace/cache"
	"ace/model"
	"ace/request"
	"ace/serializer"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BotAsk struct {
	Id      *string `form:"id" json:"id" binding:"required"`
	Content *string `form:"content" json:"content" binding:"required"`
}

func (b *BotAsk) Ask(user model.User) serializer.Response {
	var bot model.ChatBot
	if err := cache.DB.Model(&model.ChatBot{}).Where("id = ? AND user_id = ?", b.Id, user.Id).First(&bot).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		zap.L().Error("[ChatBot] Get chat bot failed", zap.Error(err))
		return serializer.DBError(err)
	}
	request.Client.Body = map[string]any{
		"db_name": bot.ChatId,
		"msg":     *b.Content,
	}
	body, err := request.Client.Post("/chat", 3)
	if err != nil {
		if err != nil {
			zap.L().Error("[ChatBot] Ask chat bot failed", zap.Error(err))
			return serializer.GeneratorError(err)
		}
	}
	return serializer.Response{
		Code: 200,
		Data: string(body),
	}
}
