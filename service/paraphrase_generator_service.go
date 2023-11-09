package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type ParaphraseGeneratorRequest struct {
	Text     string `form:"text" json:"text" binding:"required"`
	Language int    `form:"language" json:"language" binding:"required"`
}

func (t *ParaphraseGeneratorRequest) Generator(user model.User) serializer.Response {
	tone := model.OptimizedAds{
		UserId:     user.Id,
		Type:       3,
		Text:       t.Text,
		LanguageId: uint(t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.OptimizedAds{}).Create(&tone).Error; err != nil {
		tx.Rollback()
		return serializer.DBError(err)
	}

	// TODO Generate OptimizedAds Content Change Tone
	content := model.OptimizedContent{
		Type:   3,
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
