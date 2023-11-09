package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type IntroGeneratorRequest struct {
	Topic      string `form:"topic" json:"topic" binding:"required"`
	Tones      int    `form:"tones" json:"tones" binding:"required"`
	BrandVoice int    `form:"brand_voice" json:"brand_voice"`
	MinAge     int    `form:"min_age" json:"min_age"`
	MaxAge     int    `form:"max_age" json:"max_age"`
	Language   int    `form:"language" json:"language" binding:"required"`
}

func (t *IntroGeneratorRequest) Generator(user model.User) serializer.Response {
	tone := model.BlogAds{
		UserId:     user.Id,
		Type:       1,
		ToneId:     uint(t.Tones),
		VoiceId:    uint(t.BrandVoice),
		MinAge:     t.MinAge,
		MaxAge:     t.MaxAge,
		LanguageId: uint(t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.BlogAds{}).Create(&tone).Error; err != nil {
		tx.Rollback()
		return serializer.DBError(err)
	}

	// TODO Generate OptimizedAds Content Change Tone
	content := model.BlogContent{
		Type:   1,
		AdsId:  tone.Id,
		UserId: user.Id,
		Text:   "testsuite",
	}
	if err := tx.Model(&model.BlogContent{}).Create(&content).Error; err != nil {
		tx.Rollback()
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
