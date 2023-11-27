package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FreestyleHistoryIdRequest struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

func (m *FreestyleHistoryIdRequest) GetFreestyleContentById(user model.User) serializer.Response {
	var content model.EmailContent
	if err := cache.DB.Model(&model.EmailContent{}).Where("user_id = ? AND id = ?", user.Id, m.Id).First(&content).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		zap.L().Error("[Freestyle] Get email content failure", zap.Error(err))
		return serializer.DBError(err)
	}

	var email model.EmailAds
	if err := cache.DB.Model(&model.EmailAds{}).Where("id = ?", content.AdsId).First(&email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		zap.L().Error("[Freestyle] Get email ads failure", zap.Error(err))
		return serializer.DBError(err)
	}

	history := serializer.FreestyleHistory{
		Detail:     email.Detail,
		Tones:      int(email.ToneId),
		BrandVoice: int(email.VoiceId),
		Region:     int(email.Region),
		Gender:     int(email.Gender),
		MinAge:     email.MinAge,
		MaxAge:     email.MaxAge,
		Language:   int(email.LanguageId),
	}

	return history.Bind()
}
