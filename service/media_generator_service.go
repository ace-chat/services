package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type MediaGeneratorRequest struct {
	Platform    int    `form:"platform" json:"platform" binding:"required"`
	BrandName   string `form:"brand_name" json:"brand_name" binding:"required"`
	ServiceName string `form:"service_name" json:"service_name" binding:"required"`
	ServiceDesc string `form:"service_desc" json:"service_desc" binding:"required"`
	Tones       int    `form:"tones" json:"tones" binding:"required"`
	BrandVoice  int    `form:"brand_voice" json:"brand_voice"`
	Region      int    `form:"region" json:"region"`
	Gender      int    `form:"gender" json:"gender"`
	MinAge      int    `form:"min_age" json:"min_age"`
	MaxAge      int    `form:"max_age" json:"max_age"`
	Language    int    `form:"language" json:"language"`
}

func (m *MediaGeneratorRequest) Generator(user model.User) serializer.Response {
	media := model.MediaAds{
		UserId:      user.Id,
		PlatformId:  uint(m.Platform),
		BrandName:   m.BrandName,
		ServiceName: m.ServiceName,
		Description: m.ServiceDesc,
		ToneId:      uint(m.Tones),
		VoiceId:     uint(m.BrandVoice),
		Region:      uint(m.Region),
		Gender:      uint(m.Gender),
		MinAge:      m.MinAge,
		MaxAge:      m.MaxAge,
		LanguageId:  uint(m.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.MediaAds{}).Create(&media).Error; err != nil {
		tx.Rollback()
		return serializer.DBError(err)
	}

	// TODO generator social media ads content
	content := model.MediaContent{
		AdsId:  media.Id,
		UserId: user.Id,
		Text:   "testsuite",
	}
	if err := tx.Model(&model.MediaContent{}).Create(&content).Error; err != nil {
		tx.Rollback()
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
