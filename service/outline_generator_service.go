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

type OutlineGeneratorRequest struct {
	Topic      *string `form:"topic" json:"topic" binding:"required"`
	Tones      *int    `form:"tones" json:"tones" binding:"required"`
	BrandVoice *int    `form:"brand_voice" json:"brand_voice"`
	MinAge     *int    `form:"min_age" json:"min_age"`
	MaxAge     *int    `form:"max_age" json:"max_age"`
	Language   *int    `form:"language" json:"language" binding:"required"`
}

func (t *OutlineGeneratorRequest) Generator(user model.User) serializer.Response {
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
		zap.L().Error("[Outline] Get tone failure", zap.Error(err))
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
			zap.L().Error("[Outline] Get brand voice failure", zap.Error(err))
			return serializer.DBError(err)
		}
		voice = v.Content
		voiceId = v.Id
	}

	language, err := tools.GetLanguage(uint(*t.Language))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundLanguageError(err)
		}
		zap.L().Error("[Outline] Get language failure", zap.Error(err))
		return serializer.DBError(err)
	}

	blog := model.BlogAds{
		UserId:     user.Id,
		Type:       2,
		ToneId:     uint(*t.Tones),
		VoiceId:    voiceId,
		MinAge:     minimum,
		MaxAge:     maximum,
		LanguageId: uint(*t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.BlogAds{}).Create(&blog).Error; err != nil {
		zap.L().Error("[Outline] Create blog ads failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	request.Client.Body = map[string]any{
		"topic":       *t.Topic,
		"tone":        tone.Value,
		"brand_voice": voice,
		"min_age":     minimum,
		"max_age":     maximum,
		"lang":        language.Iso,
	}

	body, err := request.Client.Post(model.Url["generate_blog_outline"])
	if err != nil {
		return serializer.GeneratorError(err)
	}

	content := model.BlogContent{
		Type:   2,
		AdsId:  blog.Id,
		UserId: user.Id,
		Text:   string(body),
	}
	if err := tx.Model(&model.BlogContent{}).Create(&content).Error; err != nil {
		zap.L().Error("[Outline] Create blog content failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
