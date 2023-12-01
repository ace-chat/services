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

type FreestyleGeneratorRequest struct {
	Detail     *string `form:"detail" json:"detail" binding:"required"`
	Tones      *int    `form:"tones" json:"tones" binding:"required"`
	BrandVoice *int    `form:"brand_voice" json:"brand_voice"`
	Region     *int    `form:"region" json:"region"`
	Gender     *int    `form:"gender" json:"gender"`
	MinAge     *int    `form:"min_age" json:"min_age"`
	MaxAge     *int    `form:"max_age" json:"max_age"`
	Language   *int    `form:"language" json:"language" binding:"required"`
}

func (t *FreestyleGeneratorRequest) Generator(user model.User) serializer.Response {
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

	tone, err := tools.GetTone(uint(*t.Tones))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundToneError(err)
		}
		return serializer.DBError(err)
	}

	var voice string
	if t.BrandVoice != nil || *t.BrandVoice != 0 {
		v, err := tools.GetVoice(uint(*t.BrandVoice), user.Id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return serializer.NotFoundVoiceError(err)
			}
			return serializer.DBError(err)
		}
		voice = v.Content
	}

	var region string
	if t.Region != nil || *t.Region != 0 {
		r, err := tools.GetRegion(uint(*t.Region))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return serializer.NotFoundRegionError(err)
			}
			return serializer.DBError(err)
		}
		region = r.Country
	}

	var gender string
	if t.Gender != nil || *t.Gender != 0 {
		g, err := tools.GetGender(uint(*t.Gender))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return serializer.NotFoundGenderError(err)
			}
			return serializer.DBError(err)
		}
		gender = g.Value
	}

	language, err := tools.GetLanguage(uint(*t.Language))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundLanguageError(err)
		}
		return serializer.DBError(err)
	}

	ads := model.EmailAds{
		UserId:     user.Id,
		Type:       1,
		Detail:     *t.Detail,
		Region:     uint(*t.Region),
		Gender:     uint(*t.Gender),
		MinAge:     minimum,
		MaxAge:     maximum,
		LanguageId: uint(*t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.EmailAds{}).Create(&ads).Error; err != nil {
		zap.L().Error("[Freestyle] Create email ads failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	request.Client.Body = map[string]any{
		"text":        t.Detail,
		"tone":        tone.Value,
		"brand_voice": voice,
		"region":      region,
		"gender":      gender,
		"min_age":     minimum,
		"max_age":     maximum,
		"lang":        language.Iso,
	}

	body, err := request.Client.Post(model.Url["generate_freestyle_email_content"])
	if err != nil {
		return serializer.GeneratorError(err)
	}

	content := model.EmailContent{
		Type:   1,
		AdsId:  ads.Id,
		UserId: user.Id,
		Text:   string(body),
	}
	if err := tx.Model(&model.EmailContent{}).Create(&content).Error; err != nil {
		zap.L().Error("[Freestyle] Create email content failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
