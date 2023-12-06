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
	Platform    *int    `form:"platform" json:"platform" binding:"required"`
	BrandName   *string `form:"brand_name" json:"brand_name" binding:"required"`
	ServiceName *string `form:"service_name" json:"service_name" binding:"required"`
	ServiceDesc *string `form:"service_desc" json:"service_desc" binding:"required"`
	Tones       *int    `form:"tones" json:"tones" binding:"required"`
	BrandVoice  *int    `form:"brand_voice" json:"brand_voice"`
	Region      *int    `form:"region" json:"region"`
	Gender      *int    `form:"gender" json:"gender"`
	MinAge      *int    `form:"min_age" json:"min_age"`
	MaxAge      *int    `form:"max_age" json:"max_age"`
	Language    *int    `form:"language" json:"language" binding:"required"`
}

func (m *MediaGeneratorRequest) Generator(user model.User) serializer.Response {
	var tools utils.Common

	var minimum, maximum int
	if m.MinAge != nil {
		minimum = *m.MinAge
	} else {
		minimum = 0
	}
	if m.MaxAge != nil {
		maximum = *m.MaxAge
	} else {
		maximum = 0
	}

	platform, err := tools.GetPlatform(uint(*m.Platform))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundPlatformError(err)
		}
		zap.L().Error("[Social Media] Get platform failure", zap.Error(err))
		return serializer.DBError(err)
	}

	tone, err := tools.GetTone(uint(*m.Tones))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundToneError(err)
		}
		zap.L().Error("[Social Media] Get tone failure", zap.Error(err))
		return serializer.DBError(err)
	}

	var brandVoice string
	var brandVoiceId uint
	if m.BrandVoice != nil {
		voice, err := tools.GetVoice(uint(*m.BrandVoice), user.Id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return serializer.NotFoundVoiceError(err)
			}
			zap.L().Error("[Social Media] Get brand voice failure", zap.Error(err))
			return serializer.DBError(err)
		}
		brandVoice = voice.Content
		brandVoiceId = voice.Id
	}

	var region string
	var regionId uint
	if m.Region != nil {
		r, err := tools.GetRegion(uint(*m.Region))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return serializer.NotFoundRegionError(err)
			}
			zap.L().Error("[Social Media] Get region failure", zap.Error(err))
			return serializer.DBError(err)
		}
		region = r.Country
		regionId = r.Id
	}

	var gender string
	var genderId uint
	if m.Gender != nil {
		g, err := tools.GetGender(uint(*m.Gender))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return serializer.NotFoundRegionError(err)
			}
			zap.L().Error("[Social Media] Get gender failure", zap.Error(err))
			return serializer.DBError(err)
		}
		gender = g.Value
		genderId = g.Id
	}

	language, err := tools.GetLanguage(uint(*m.Language))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundLanguageError(err)
		}
		zap.L().Error("[Social Media] Get language failure", zap.Error(err))
		return serializer.DBError(err)
	}

	media := model.MediaAds{
		UserId:      user.Id,
		PlatformId:  uint(*m.Platform),
		BrandName:   *m.BrandName,
		ServiceName: *m.ServiceName,
		Description: *m.ServiceDesc,
		ToneId:      uint(*m.Tones),
		VoiceId:     brandVoiceId,
		Region:      regionId,
		Gender:      genderId,
		MinAge:      minimum,
		MaxAge:      maximum,
		LanguageId:  uint(*m.Language),
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
		"brand_voice":         brandVoice,
		"region":              region,
		"gender":              gender,
		"min_age":             minimum,
		"max_age":             maximum,
		"lang":                language.Iso,
	}

	body, err := request.Client.Post(model.Url["generator_social_media"])
	if err != nil {
		return serializer.GeneratorError(err)
	}

	zap.L().Info("[Generate] Social Media Content", zap.String("content", string(body)))

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
