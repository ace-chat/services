package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MediaHistoryIdRequest struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

func (m *MediaHistoryIdRequest) GetMediaContentById(user model.User) serializer.Response {
	var content model.MediaContent
	if err := cache.DB.Model(&model.MediaContent{}).Where("user_id = ? AND id = ?", user.Id, m.Id).First(&content).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		zap.L().Error("[Media] Get social media ads content failed", zap.Error(err))
		return serializer.DBError(err)
	}

	var media model.MediaAds
	if err := cache.DB.Model(&model.MediaAds{}).Where("id = ?", content.AdsId).First(&media).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		zap.L().Error("[Media] Get social media ads failed", zap.Error(err))
		return serializer.DBError(err)
	}

	history := serializer.MediaAndEngineHistory{
		Platform:    int(media.PlatformId),
		BrandName:   media.BrandName,
		ServiceName: media.ServiceName,
		ServiceDesc: media.Description,
		Tones:       int(media.ToneId),
		BrandVoice:  int(media.VoiceId),
		Region:      int(media.Region),
		Gender:      int(media.Gender),
		MinAge:      media.MinAge,
		MaxAge:      media.MaxAge,
		Language:    int(media.LanguageId),
		Content:     content.Text,
	}

	return history.Bind()
}
