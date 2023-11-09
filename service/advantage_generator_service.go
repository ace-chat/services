package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type AdvantageGeneratorRequest struct {
	BrandName   string `form:"brand_name" json:"brand_name" binding:"required"`
	ServiceName string `form:"service_name" json:"service_name"`
	BrandDesc   string `form:"brand_desc" json:"brand_desc"`
	Tones       int    `form:"tones" json:"tones" binding:"required"`
	BrandVoice  int    `form:"brand_voice" json:"brand_voice"`
	Region      int    `form:"region" json:"region"`
	Gender      int    `form:"gender" json:"gender"`
	MinAge      int    `form:"min_age" json:"min_age"`
	MaxAge      int    `form:"max_age" json:"max_age"`
	Language    int    `form:"language" json:"language" binding:"required"`
}

func (t *AdvantageGeneratorRequest) Generator(user model.User) serializer.Response {
	tone := model.EmailAds{
		UserId:      user.Id,
		Type:        4,
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
	if err := tx.Model(&model.EmailAds{}).Create(&tone).Error; err != nil {
		tx.Rollback()
		return serializer.DBError(err)
	}

	// TODO Generate OptimizedAds Content Change Tone
	content := model.EmailContent{
		Type:   4,
		AdsId:  tone.Id,
		UserId: user.Id,
		Text:   "testsuite",
	}
	if err := tx.Model(&model.EmailContent{}).Create(&content).Error; err != nil {
		tx.Rollback()
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
