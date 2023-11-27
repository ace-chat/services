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
	Topic      string `form:"topic" json:"topic" binding:"required"`
	Tones      int    `form:"tones" json:"tones" binding:"required"`
	BrandVoice int    `form:"brand_voice" json:"brand_voice"`
	MinAge     int    `form:"min_age" json:"min_age"`
	MaxAge     int    `form:"max_age" json:"max_age"`
	Language   int    `form:"language" json:"language" binding:"required"`
}

func (t *OutlineGeneratorRequest) Generator(user model.User) serializer.Response {
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

	language, err := tools.GetLanguage(uint(t.Language))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundLanguageError(err)
		}
		return serializer.DBError(err)
	}

	blog := model.BlogAds{
		UserId:     user.Id,
		Type:       2,
		ToneId:     uint(t.Tones),
		VoiceId:    uint(t.BrandVoice),
		MinAge:     t.MinAge,
		MaxAge:     t.MaxAge,
		LanguageId: uint(t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.BlogAds{}).Create(&blog).Error; err != nil {
		zap.L().Error("[Outline] Create blog ads failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	request.Client.Body = map[string]any{
		"topic":       t.Topic,
		"tone":        tone.Value,
		"brand_voice": voice.Content,
		"min_age":     t.MinAge,
		"max_age":     t.MaxAge,
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
