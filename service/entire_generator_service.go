package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type EntireGeneratorRequest struct {
	Topic        string `form:"topic" json:"topic" binding:"required"`
	Tones        int    `form:"tones" json:"tones" binding:"required"`
	BrandVoice   int    `form:"brand_voice" json:"brand_voice"`
	Keyword      string `form:"keyword" json:"keyword"`
	MinAge       int    `form:"min_age" json:"min_age"`
	MaxAge       int    `form:"max_age" json:"max_age"`
	WordCount    int    `form:"word_count" json:"word_count"`
	OtherDetails string `form:"other_details" json:"other_details"`
	Language     int    `form:"language" json:"language" binding:"required"`
}

func (t *EntireGeneratorRequest) Generator(user model.User) serializer.Response {
	tone := model.BlogAds{
		UserId:       user.Id,
		Type:         3,
		ToneId:       uint(t.Tones),
		VoiceId:      uint(t.BrandVoice),
		Keyword:      t.Keyword,
		MinAge:       t.MinAge,
		MaxAge:       t.MaxAge,
		WordCount:    t.WordCount,
		OtherDetails: t.OtherDetails,
		LanguageId:   uint(t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.BlogAds{}).Create(&tone).Error; err != nil {
		tx.Rollback()
		return serializer.DBError(err)
	}

	// TODO Generate OptimizedAds Content Change Tone
	content := model.BlogContent{
		Type:   3,
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
