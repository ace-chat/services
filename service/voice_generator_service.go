package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type VoiceGeneratorRequest struct {
	Text       string `form:"text" json:"text" binding:"required"`
	BrandVoice int    `form:"brand_voice" json:"brand_voice" binding:"required"`
	Language   int    `form:"language" json:"language" binding:"required"`
}

func (t *VoiceGeneratorRequest) Generator(user model.User) serializer.Response {
	tone := model.OptimizedAds{
		UserId:     user.Id,
		Type:       4,
		Text:       t.Text,
		VoiceId:    uint(t.BrandVoice),
		LanguageId: uint(t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.OptimizedAds{}).Create(&tone).Error; err != nil {
		tx.Rollback()
		return serializer.DBError(err)
	}

	// TODO Generate OptimizedAds Content Change Tone
	content := model.OptimizedContent{
		Type:   4,
		AdsId:  tone.Id,
		UserId: user.Id,
		Text:   "testsuite",
	}
	if err := tx.Model(&model.OptimizedContent{}).Create(&content).Error; err != nil {
		tx.Rollback()
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
