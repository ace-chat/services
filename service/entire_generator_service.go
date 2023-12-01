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

type EntireGeneratorRequest struct {
	Topic        *string `form:"topic" json:"topic" binding:"required"`
	Tones        *int    `form:"tones" json:"tones" binding:"required"`
	Type         *int    `form:"type" json:"type" binding:"required"`
	BrandVoice   *int    `form:"brand_voice" json:"brand_voice"`
	Keyword      *string `form:"keyword" json:"keyword"`
	MinAge       *int    `form:"min_age" json:"min_age"`
	MaxAge       *int    `form:"max_age" json:"max_age"`
	WordCount    *int    `form:"word_count" json:"word_count" binding:"required"`
	OtherDetails *string `form:"other_details" json:"other_details"`
	Language     *int    `form:"language" json:"language" binding:"required"`
}

func (t *EntireGeneratorRequest) Generator(user model.User) serializer.Response {
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

	var keyword string
	if t.Keyword != nil {
		keyword = *t.Keyword
	} else {
		keyword = ""
	}

	var detail string
	if t.OtherDetails != nil {
		detail = *t.OtherDetails
	} else {
		detail = ""
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

	language, err := tools.GetLanguage(uint(*t.Language))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundLanguageError(err)
		}
		return serializer.DBError(err)
	}

	ty, err := tools.GetType(uint(*t.Type))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundTypeError(err)
		}
		return serializer.DBError(err)
	}

	blog := model.BlogAds{
		UserId:       user.Id,
		Type:         3,
		BlogType:     uint(*t.Type),
		ToneId:       uint(*t.Tones),
		VoiceId:      uint(*t.BrandVoice),
		Keyword:      keyword,
		MinAge:       minimum,
		MaxAge:       maximum,
		WordCount:    *t.WordCount,
		OtherDetails: detail,
		LanguageId:   uint(*t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.BlogAds{}).Create(&blog).Error; err != nil {
		zap.L().Error("[Entire] Create blog ads failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	request.Client.Body = map[string]any{
		"topic":         t.Topic,
		"tone":          tone.Value,
		"brand_voice":   voice,
		"keywords":      keyword,
		"min_age":       minimum,
		"max_age":       maximum,
		"word_count":    t.WordCount,
		"other_details": detail,
		"lang":          language.Name,
		"type":          ty.Value,
	}

	body, err := request.Client.Post(model.Url["generate_blog_entire"])
	if err != nil {
		return serializer.GeneratorError(err)
	}

	content := model.BlogContent{
		Type:   3,
		AdsId:  blog.Id,
		UserId: user.Id,
		Text:   string(body),
	}
	if err := tx.Model(&model.BlogContent{}).Create(&content).Error; err != nil {
		zap.L().Error("[Entire] Create blog content failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
