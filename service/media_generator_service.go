package service

import (
	"ace/cache"
	"ace/model"
	"ace/request"
	"ace/serializer"
	"ace/utils"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

type MediaGeneratorRequest struct {
	Platform    int    `form:"platform" json:"platform" binding:"required"`
	BrandName   string `form:"brand_name" json:"brand_name" binding:"required"`
	ServiceName string `form:"service_name" json:"service_name" binding:"required"`
	ServiceDesc string `form:"service_desc" json:"service_desc" binding:"required"`
	Tones       int    `form:"tones" json:"tones" binding:"required"`
	BrandVoice  int    `form:"brand_voice" json:"brand_voice"`
	Region      int    `form:"region" json:"region"`
	Gender      int    `form:"gender" json:"gender"`
	MinAge      int    `form:"min_age" json:"min_age"`
	MaxAge      int    `form:"max_age" json:"max_age"`
	Language    int    `form:"language" json:"language"`
}

func (m *MediaGeneratorRequest) Generator(user model.User) serializer.Response {
	var tools utils.Common

	platform, err := tools.GetPlatform(uint(m.Platform))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundPlatformError(err)
		}
		return serializer.DBError(err)
	}

	tone, err := tools.GetTone(uint(m.Tones))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundToneError(err)
		}
		return serializer.DBError(err)
	}

	voice, err := tools.GetVoice(uint(m.BrandVoice), user.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundVoiceError(err)
		}
		return serializer.DBError(err)
	}

	region, err := tools.GetRegion(uint(m.Region))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundRegionError(err)
		}
		return serializer.DBError(err)
	}

	gender, err := tools.GetGender(uint(m.Gender))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundRegionError(err)
		}
		return serializer.DBError(err)
	}

	language, err := tools.GetLanguage(uint(m.Language))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundLanguageError(err)
		}
		return serializer.DBError(err)
	}

	media := model.MediaAds{
		UserId:      user.Id,
		PlatformId:  uint(m.Platform),
		BrandName:   m.BrandName,
		ServiceName: m.ServiceName,
		Description: m.ServiceDesc,
		ToneId:      uint(m.Tones),
		VoiceId:     uint(m.BrandVoice),
		Region:      uint(m.Region),
		Gender:      uint(m.Gender),
		MinAge:      m.MinAge,
		MaxAge:      m.MaxAge,
		LanguageId:  uint(m.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.MediaAds{}).Create(&media).Error; err != nil {
		zap.L().Error("[Media] Create social media ads failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	request.Client.Body = map[string]any{
		"platform":            platform.Value,
		"brand_name":          m.BrandName,
		"product_name":        m.ServiceName,
		"product_description": m.ServiceDesc,
		"tone":                tone.Value,
		"brand_voice":         voice.Content,
		"region":              region.Iso,
		"gender":              gender.Value,
		"min_age":             m.MinAge,
		"max_age":             m.MaxAge,
		"lang":                language.Iso,
	}

	body, err := request.Client.Post(model.Url["generator_social_media"])
	if err != nil {
		return serializer.GeneratorError(err)
	}

	content := model.MediaContent{
		AdsId:  media.Id,
		UserId: user.Id,
		Text:   string(body),
	}
	if err := tx.Model(&model.MediaContent{}).Create(&content).Error; err != nil {
		zap.L().Error("[Media] Create social media ads content failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
