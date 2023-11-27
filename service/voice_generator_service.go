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

type VoiceGeneratorRequest struct {
	Text       string `form:"text" json:"text" binding:"required"`
	BrandVoice int    `form:"brand_voice" json:"brand_voice" binding:"required"`
	Language   int    `form:"language" json:"language" binding:"required"`
}

func (t *VoiceGeneratorRequest) Generator(user model.User) serializer.Response {
	var tools utils.Common

	voice, err := tools.GetVoice(uint(t.BrandVoice), user.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundVoiceError(err)
		}
		return serializer.DBError(err)
	}

	language, err := tools.GetLanguage(uint(t.Language))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundLanguageError(err)
		}
		return serializer.DBError(err)
	}

	ads := model.OptimizedAds{
		UserId:     user.Id,
		Type:       4,
		Text:       t.Text,
		VoiceId:    uint(t.BrandVoice),
		LanguageId: uint(t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.OptimizedAds{}).Create(&ads).Error; err != nil {
		zap.L().Error("[Voice] Create optimized ads failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	request.Client.Body = map[string]any{
		"text":        t.Text,
		"brand_voice": voice.Content,
		"lang":        language.Iso,
	}

	body, err := request.Client.Post(model.Url["generate_optimize_match_brand_voice"])
	if err != nil {
		return serializer.GeneratorError(err)
	}

	zap.L().Info("[Generate] Optimize Match Brand Voice Content", zap.String("content", string(body)))

	content := model.OptimizedContent{
		Type:   4,
		AdsId:  ads.Id,
		UserId: user.Id,
		Text:   string(body),
	}
	if err := tx.Model(&model.OptimizedContent{}).Create(&content).Error; err != nil {
		zap.L().Error("[Voice] Create optimized content failure", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
