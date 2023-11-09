package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type SummarizeGeneratorRequest struct {
	Text      string `form:"text" json:"text" binding:"required"`
	WordCount int    `form:"word_count" json:"word_count" binding:"required"`
	Language  int    `form:"language" json:"language" binding:"required"`
}

func (t *SummarizeGeneratorRequest) Generator(user model.User) serializer.Response {
	tone := model.OptimizedAds{
		UserId:     user.Id,
		Type:       2,
		Text:       t.Text,
		WordCount:  t.WordCount,
		LanguageId: uint(t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.OptimizedAds{}).Create(&tone).Error; err != nil {
		tx.Rollback()
		return serializer.DBError(err)
	}

	// TODO Generate OptimizedAds Content Change Tone
	content := model.OptimizedContent{
		Type:   2,
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
