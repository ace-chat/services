package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"errors"
	"gorm.io/gorm"
)

type AdvantageHistoryIdRequest struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

func (m *AdvantageHistoryIdRequest) GetToneContentById(user model.User) serializer.Response {
	var content model.EmailContent
	if err := cache.DB.Model(&model.EmailContent{}).Where("user_id = ? AND id = ?", user.Id, m.Id).First(&content).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		return serializer.DBError(err)
	}

	var email model.EmailAds
	if err := cache.DB.Model(&model.EmailAds{}).Where("id = ?", content.AdsId).First(&email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		return serializer.DBError(err)
	}

	history := serializer.MarketingAndWelcomeAndAdvantageHistory{
		BrandName:   email.BrandName,
		ServiceName: email.ServiceName,
		BrandDesc:   email.BrandDesc,
		Tones:       int(email.ToneId),
		BrandVoice:  int(email.VoiceId),
		Region:      int(email.Region),
		Gender:      int(email.Gender),
		MinAge:      email.MinAge,
		MaxAge:      email.MaxAge,
		Language:    int(email.LanguageId),
	}

	return history.Bind()
}
