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
	BrandName   *string `form:"brand_name" json:"brand_name" binding:"required"`
	BrandDesc   *string `form:"brand_desc" json:"brand_desc"`
	ServiceName *string `form:"service_name" json:"service_name"`
	ServiceDesc *string `form:"service_desc" json:"service_desc"`
	Tones       *int    `form:"tones" json:"tones" binding:"required"`
	BrandVoice  *int    `form:"brand_voice" json:"brand_voice"`
	Region      *int    `form:"region" json:"region"`
	Gender      *int    `form:"gender" json:"gender"`
	MinAge      *int    `form:"min_age" json:"min_age"`
	MaxAge      *int    `form:"max_age" json:"max_age"`
	Language    *int    `form:"language" json:"language" binding:"required"`
}

func (t *WelcomeGeneratorRequest) Generator(user model.User) serializer.Response {
	var tools utils.Common

	var minimum, maximum int
	if t.MinAge != nil {
		minimum = *t.MinAge
	} else {
		minimum = 0
	}
	if t.MaxAge != nil {
		maximum = *t.MaxAge
	} else {
		maximum = 0
	}

	var brandDesc, serviceName, serviceDesc string
	if t.BrandDesc != nil {
		brandDesc = *t.BrandDesc
	} else {
		brandDesc = ""
	}
	if t.ServiceName != nil {
		serviceName = *t.ServiceName
	} else {
		serviceName = ""
	}
	if t.ServiceDesc != nil {
		serviceDesc = *t.ServiceDesc
	} else {
		serviceDesc = ""
	}

	tone, err := tools.GetTone(uint(*t.Tones))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundToneError(err)
		}
		zap.L().Error("[Welcome] Get tone failure", zap.Error(err))
		return serializer.DBError(err)
	}

	var voice string
	var voiceId uint
	if t.BrandVoice != nil {
		v, err := tools.GetVoice(uint(*t.BrandVoice), user.Id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return serializer.NotFoundVoiceError(err)
			}
			zap.L().Error("[Welcome] Get brand voice failure", zap.Error(err))
			return serializer.DBError(err)
		}
		voice = v.Content
		voiceId = v.Id
	}

	var region string
	var regionId uint
	if t.Region != nil {
		r, err := tools.GetRegion(uint(*t.Region))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return serializer.NotFoundRegionError(err)
			}
			zap.L().Error("[Welcome] Get region failure", zap.Error(err))
			return serializer.DBError(err)
		}
		region = r.Country
		regionId = r.Id
	}

	var gender string
	var genderId uint
	if t.Gender != nil {
		g, err := tools.GetGender(uint(*t.Gender))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return serializer.NotFoundGenderError(err)
			}
			zap.L().Error("[Welcome] Get gender failure", zap.Error(err))
			return serializer.DBError(err)
		}
		gender = g.Value
		genderId = g.Id
	}

	language, err := tools.GetLanguage(uint(*t.Language))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundLanguageError(err)
		}
		zap.L().Error("[Welcome] Get language failure", zap.Error(err))
		return serializer.DBError(err)
	}

	ads := model.EmailAds{
		UserId:      user.Id,
		Type:        3,
		BrandName:   *t.BrandName,
		ServiceName: serviceName,
		ServiceDesc: serviceDesc,
		BrandDesc:   brandDesc,
		ToneId:      uint(*t.Tones),
		VoiceId:     voiceId,
		Region:      regionId,
		Gender:      genderId,
		MinAge:      minimum,
		MaxAge:      maximum,
		LanguageId:  uint(*t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.EmailAds{}).Create(&ads).Error; err != nil {
		zap.L().Error("[Welcome] Create email ads failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	request.Client.Body = map[string]any{
		"brand_name":          *t.BrandName,
		"brand_description":   brandDesc,
		"product_name":        serviceName,
		"product_description": serviceDesc,
		"tone":                tone.Value,
		"brand_voice":         voice,
		"region":              region,
		"gender":              gender,
		"min_age":             minimum,
		"max_age":             maximum,
		"lang":                language.Iso,
	}

	body, err := request.Client.Post(model.Url["generate_welcome_email_content"], false)
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
