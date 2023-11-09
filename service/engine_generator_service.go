package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type EngineGeneratorRequest struct {
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

func (e *EngineGeneratorRequest) Generator(user model.User) serializer.Response {
	media := model.EngineAds{
		UserId:      user.Id,
		PlatformId:  uint(e.Platform),
		BrandName:   e.BrandName,
		ServiceName: e.ServiceName,
		Description: e.ServiceDesc,
		ToneId:      uint(e.Tones),
		VoiceId:     uint(e.BrandVoice),
		Region:      uint(e.Region),
		Gender:      uint(e.Gender),
		MinAge:      e.MinAge,
		MaxAge:      e.MaxAge,
		LanguageId:  uint(e.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.EngineAds{}).Create(&media).Error; err != nil {
		tx.Rollback()
		return serializer.DBError(err)
	}

	// TODO generator ai search engine ads content
	content := model.EngineContent{
		AdsId:  media.Id,
		UserId: user.Id,
		Text:   "testsuite",
	}
	if err := tx.Model(&model.EngineContent{}).Create(&content).Error; err != nil {
		tx.Rollback()
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
