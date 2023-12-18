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

type EngineGeneratorRequest struct {
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

func (e *EngineGeneratorRequest) Generator(user model.User) serializer.Response {
	var tools utils.Common

	var minimum, maximum int
	if e.MinAge != nil {
		minimum = *e.MinAge
	} else {
		minimum = 0
	}
	if e.MaxAge != nil {
		maximum = *e.MaxAge
	} else {
		maximum = 0
	}

	platform, err := tools.GetPlatform(uint(*e.Platform))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundPlatformError(err)
		}
		zap.L().Error("[Search Engine] Get platform failed", zap.Error(err))
		return serializer.DBError(err)
	}

	tone, err := tools.GetTone(uint(*e.Tones))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundToneError(err)
		}
		zap.L().Error("[Search Engine] Get tone failed", zap.Error(err))
		return serializer.DBError(err)
	}

	var brandVoice string
	var brandVoiceId uint
	if e.BrandVoice != nil {
		voice, err := tools.GetVoice(uint(*e.BrandVoice), user.Id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return serializer.NotFoundVoiceError(err)
			}
			zap.L().Error("[Search Engine] Get brand voice failed", zap.Error(err))
			return serializer.DBError(err)
		}
		brandVoice = voice.Content
		brandVoiceId = voice.Id
	}

	var region string
	var regionId uint
	if e.Region != nil {
		r, err := tools.GetRegion(uint(*e.Region))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return serializer.NotFoundRegionError(err)
			}
			zap.L().Error("[Search Engine] Get region failed", zap.Error(err))
			return serializer.DBError(err)
		}
		region = r.Country
		regionId = r.Id
	}

	var gender string
	var genderId uint
	if e.Gender != nil {
		g, err := tools.GetGender(uint(*e.Gender))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return serializer.NotFoundRegionError(err)
			}
			zap.L().Error("[Search Engine] Get gender failed", zap.Error(err))
			return serializer.DBError(err)
		}
		gender = g.Value
		genderId = g.Id
	}

	language, err := tools.GetLanguage(uint(*e.Language))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundLanguageError(err)
		}
		zap.L().Error("[Search Engine] Get language failed", zap.Error(err))
		return serializer.DBError(err)
	}

	media := model.EngineAds{
		UserId:      user.Id,
		PlatformId:  uint(*e.Platform),
		BrandName:   *e.BrandName,
		ServiceName: *e.ServiceName,
		Description: *e.ServiceDesc,
		ToneId:      uint(*e.Tones),
		VoiceId:     brandVoiceId,
		Region:      regionId,
		Gender:      genderId,
		MinAge:      minimum,
		MaxAge:      maximum,
		LanguageId:  uint(*e.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.EngineAds{}).Create(&media).Error; err != nil {
		zap.L().Error("[Engine] Create search engine ads failed", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	request.Client.Body = map[string]any{
		"platform":            platform.Value,
		"brand_name":          *e.BrandName,
		"product_name":        *e.ServiceName,
		"product_description": *e.ServiceDesc,
		"tone":                tone.Value,
		"brand_voice":         brandVoice,
		"region":              region,
		"gender":              gender,
		"min_age":             minimum,
		"max_age":             maximum,
		"lang":                language.Iso,
	}

	body, err := request.Client.Post(model.Url["generator_search_engine"], false)
	if err != nil {
		tx.Rollback()
		return serializer.GeneratorError(err)
	}

	zap.L().Info("[Generate] Search Engine Content", zap.String("content", string(body)))

	content := model.EngineContent{
		AdsId:  media.Id,
		UserId: user.Id,
		Text:   string(body),
	}
	if err := tx.Model(&model.EngineContent{}).Create(&content).Error; err != nil {
		tx.Rollback()
		zap.L().Error("[Engine] Create search engine ads content failed", zap.Error(err))
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
