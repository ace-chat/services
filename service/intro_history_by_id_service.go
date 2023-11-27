package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IntroHistoryIdRequest struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

func (m *IntroHistoryIdRequest) GetIntroContentById(user model.User) serializer.Response {
	var content model.BlogContent
	if err := cache.DB.Model(&model.BlogContent{}).Where("user_id = ? AND id = ?", user.Id, m.Id).First(&content).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		zap.L().Error("[Intro] Get blog content failure", zap.Error(err))
		return serializer.DBError(err)
	}

	var blog model.BlogAds
	if err := cache.DB.Model(&model.BlogAds{}).Where("id = ?", content.AdsId).First(&blog).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		zap.L().Error("[Intro] Get blog ads failure", zap.Error(err))
		return serializer.DBError(err)
	}

	history := serializer.IntroAndOutlineHistory{
		Topic:      blog.Topic,
		Tones:      int(blog.ToneId),
		BrandVoice: int(blog.VoiceId),
		MinAge:     blog.MinAge,
		MaxAge:     blog.MaxAge,
		Language:   int(blog.LanguageId),
	}

	return history.Bind()
}
