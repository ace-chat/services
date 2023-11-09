package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type AudienceGeneratorRequest struct {
	Text     string `form:"text" json:"text" binding:"required"`
	Region   int    `form:"region" json:"region"`
	Gender   int    `form:"gender" json:"gender"`
	MinAge   int    `form:"min_age" json:"min_age"`
	MaxAge   int    `form:"max_age" json:"max_age"`
	Language int    `form:"language" json:"language" binding:"required"`
}

func (t *AudienceGeneratorRequest) Generator(user model.User) serializer.Response {
	tone := model.OptimizedAds{
		UserId:     user.Id,
		Type:       5,
		Text:       t.Text,
		Region:     uint(t.Region),
		Gender:     uint(t.Gender),
		MinAge:     t.MinAge,
		MaxAge:     t.MaxAge,
		LanguageId: uint(t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.OptimizedAds{}).Create(&tone).Error; err != nil {
		tx.Rollback()
		return serializer.DBError(err)
	}

	// TODO Generate OptimizedAds Content Change Tone
	content := model.OptimizedContent{
		Type:   5,
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
