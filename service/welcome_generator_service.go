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

type WelcomeGeneratorRequest struct {
	BrandName   string `form:"brand_name" json:"brand_name" binding:"required"`
	BrandDesc   string `form:"brand_desc" json:"brand_desc"`
	ServiceName string `form:"service_name" json:"service_name"`
	ServiceDesc string `form:"service_desc" json:"service_desc"`
	Tones       int    `form:"tones" json:"tones" binding:"required"`
	BrandVoice  int    `form:"brand_voice" json:"brand_voice"`
	Region      int    `form:"region" json:"region"`
	Gender      int    `form:"gender" json:"gender"`
	MinAge      int    `form:"min_age" json:"min_age"`
	MaxAge      int    `form:"max_age" json:"max_age"`
	Language    int    `form:"language" json:"language" binding:"required"`
}

func (t *WelcomeGeneratorRequest) Generator(user model.User) serializer.Response {
	var tools utils.Common

	tone, err := tools.GetTone(uint(t.Tones))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundToneError(err)
		}
		return serializer.DBError(err)
	}

	voice, err := tools.GetVoice(uint(t.BrandVoice), user.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundVoiceError(err)
		}
		return serializer.DBError(err)
	}

	region, err := tools.GetRegion(uint(t.Region))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundRegionError(err)
		}
		return serializer.DBError(err)
	}

	gender, err := tools.GetGender(uint(t.Gender))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundGenderError(err)
		}
		return serializer.DBError(err)
	}

	language, err := tools.GetLanguage(uint(t.Language))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundLanguageError(err)
		}
		return serializer.DBError(err)
	}

	ads := model.EmailAds{
		UserId:      user.Id,
		Type:        3,
		BrandName:   t.BrandName,
		ServiceName: t.ServiceName,
		BrandDesc:   t.BrandDesc,
		ToneId:      uint(t.Tones),
		VoiceId:     uint(t.BrandVoice),
		Region:      uint(t.Region),
		Gender:      uint(t.Gender),
		MinAge:      t.MinAge,
		MaxAge:      t.MaxAge,
		LanguageId:  uint(t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.EmailAds{}).Create(&ads).Error; err != nil {
		zap.L().Error("[Welcome] Create email ads failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	request.Client.Body = map[string]any{
		"brand_name":          t.BrandName,
		"brand_description":   t.BrandDesc,
		"product_name":        t.ServiceName,
		"product_description": t.ServiceDesc,
		"tone":                tone.Value,
		"brand_voice":         voice.Content,
		"region":              region.Iso,
		"gender":              gender.Value,
		"min_age":             t.MinAge,
		"max_age":             t.MaxAge,
		"lang":                language.Iso,
	}

	body, err := request.Client.Post(model.Url["generate_welcome_email_content"])
	if err != nil {
		return serializer.GeneratorError(err)
	}

	content := model.EmailContent{
		Type:   3,
		AdsId:  ads.Id,
		UserId: user.Id,
		Text:   string(body),
	}
	if err := tx.Model(&model.EmailContent{}).Create(&content).Error; err != nil {
		zap.L().Error("[Welcome] Create email content failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
