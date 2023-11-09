package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"errors"
	"gorm.io/gorm"
)

type ToneHistoryIdRequest struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

func (m *ToneHistoryIdRequest) GetToneContentById(user model.User) serializer.Response {
	var content model.OptimizedContent
	if err := cache.DB.Model(&model.OptimizedContent{}).Where("user_id = ? AND id = ?", user.Id, m.Id).First(&content).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		return serializer.DBError(err)
	}

	var optimized model.OptimizedAds
	if err := cache.DB.Model(&model.OptimizedAds{}).Where("id = ?", content.AdsId).First(&optimized).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		return serializer.DBError(err)
	}

	history := serializer.ToneHistory{
		Text:     optimized.Text,
		Tones:    int(optimized.ToneId),
		Language: int(optimized.LanguageId),
	}

	return history.Bind()
}
