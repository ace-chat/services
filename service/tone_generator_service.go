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

type ToneGeneratorRequest struct {
	Text     *string `form:"text" json:"text" binding:"required"`
	Tones    *int    `form:"tones" json:"tones" binding:"required"`
	Language *int    `form:"language" json:"language" binding:"required"`
}

func (t *ToneGeneratorRequest) Generator(user model.User) serializer.Response {
	var tools utils.Common

	tone, err := tools.GetTone(uint(*t.Tones))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundToneError(err)
		}
		return serializer.DBError(err)
	}

	language, err := tools.GetLanguage(uint(*t.Language))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundLanguageError(err)
		}
		return serializer.DBError(err)
	}

	ads := model.OptimizedAds{
		UserId:     user.Id,
		Type:       1,
		Text:       *t.Text,
		ToneId:     uint(*t.Tones),
		LanguageId: uint(*t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.OptimizedAds{}).Create(&ads).Error; err != nil {
		zap.L().Error("[Tone] Create change tone failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	request.Client.Body = map[string]any{
		"text": *t.Text,
		"tone": tone.Value,
		"lang": language.Name,
	}

	body, err := request.Client.Post(model.Url["generate_optimize_change_tone"], false)
	if err != nil {
		tx.Rollback()
		return serializer.GeneratorError(err)
	}

	zap.L().Info("[Generate] Optimize Change Tone Content", zap.String("content", string(body)))

	content := model.OptimizedContent{
		Type:   1,
		AdsId:  ads.Id,
		UserId: user.Id,
		Text:   string(body),
	}
	if err := tx.Model(&model.OptimizedContent{}).Create(&content).Error; err != nil {
		zap.L().Error("[Tone] Create optimized content failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
