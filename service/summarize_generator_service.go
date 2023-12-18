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

type SummarizeGeneratorRequest struct {
	Text      *string `form:"text" json:"text" binding:"required"`
	WordCount *int    `form:"word_count" json:"word_count" binding:"required"`
	Language  *int    `form:"language" json:"language" binding:"required"`
}

func (t *SummarizeGeneratorRequest) Generator(user model.User) serializer.Response {
	var tools utils.Common

	language, err := tools.GetLanguage(uint(*t.Language))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundLanguageError(err)
		}
		return serializer.DBError(err)
	}

	ads := model.OptimizedAds{
		UserId:     user.Id,
		Type:       2,
		Text:       *t.Text,
		WordCount:  *t.WordCount,
		LanguageId: uint(*t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.OptimizedAds{}).Create(&ads).Error; err != nil {
		zap.L().Error("[Summarize] Create optimized ads failed", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	request.Client.Body = map[string]any{
		"text":       *t.Text,
		"word_count": *t.WordCount,
		"lang":       language.Name,
	}
	body, err := request.Client.Post(model.Url["generate_optimize_summarize"], false)
	if err != nil {
		tx.Rollback()
		return serializer.GeneratorError(err)
	}

	zap.L().Info("[Generate] Optimize Summarize Content", zap.String("content", string(body)))

	content := model.OptimizedContent{
		Type:   2,
		AdsId:  ads.Id,
		UserId: user.Id,
		Text:   string(body),
	}
	if err := tx.Model(&model.OptimizedContent{}).Create(&content).Error; err != nil {
		zap.L().Error("[Summarize] Create optimized ads content failed", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
