package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	histories := make([]model.ChatHistory, 0)
	ctx := context.Background()
	cursor, err := cache.Mongo.Collection(bot.ChatId).Find(ctx, bson.D{{}})
	if err != nil {
		if errors.Is(mongo.ErrNilDocument, err) {
			return serializer.Response{
				Code: 200,
				Data: histories,
			}
		}
		return serializer.MongoError(err)
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &histories); err != nil {
		zap.L().Error("[ChatBot] Parse history failed", zap.Error(err))
	}

	h := make([]model.History, 0)
	for _, history := range histories {
		var subHistory model.ChatSubHistory
		err := json.Unmarshal([]byte(history.History), &subHistory)
		if err != nil {
			zap.L().Error("[ChatBot] Unmarshal history failed", zap.Error(err))
			continue
		}
		ht := model.History{
			Type:    subHistory.Type,
			Content: subHistory.Data.Content,
			Time:    history.SessionId,
		}
		h = append(h, ht)
	}

	return serializer.Response{
		Code: 200,
		Data: h,
	}
}
