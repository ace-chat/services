package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AudienceHistoryIdRequest struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

func (m *AudienceHistoryIdRequest) GetAudienceContentById(user model.User) serializer.Response {
	var content model.OptimizedContent
	if err := cache.DB.Model(&model.OptimizedContent{}).Where("user_id = ? AND id = ?", user.Id, m.Id).First(&content).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		zap.L().Error("[Audience] Get optimized content failure", zap.Error(err))
		return serializer.DBError(err)
	}

	var optimized model.OptimizedAds
	if err := cache.DB.Model(&model.OptimizedAds{}).Where("id = ?", content.AdsId).First(&optimized).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundError(err)
		}
		zap.L().Error("[Audience] Get optimized ads failure", zap.Error(err))
		return serializer.DBError(err)
	}

	history := serializer.AudienceHistory{
		Text:     optimized.Text,
		Region:   int(optimized.Region),
		Gender:   int(optimized.Gender),
		MinAge:   optimized.MinAge,
		MaxAge:   optimized.MaxAge,
		Language: int(optimized.LanguageId),
	}

	return history.Bind()
}
