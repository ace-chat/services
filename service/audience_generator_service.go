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

type AudienceGeneratorRequest struct {
	Text     *string `form:"text" json:"text" binding:"required"`
	Region   *int    `form:"region" json:"region"`
	Gender   *int    `form:"gender" json:"gender"`
	MinAge   *int    `form:"min_age" json:"min_age"`
	MaxAge   *int    `form:"max_age" json:"max_age"`
	Language *int    `form:"language" json:"language" binding:"required"`
}

func (t *AudienceGeneratorRequest) Generator(user model.User) serializer.Response {
	var tools utils.Common

	var minimum, maximum int
	if t.MinAge != nil {
		minimum = *t.MinAge
	} else {
		minimum = 0
	}
	if t.MaxAge != nil {
		maximum = *t.MaxAge
	} else {
		maximum = 0
	}

	var region string
	var regionId uint
	if t.Region != nil {
		r, err := tools.GetRegion(uint(*t.Region))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return serializer.NotFoundRegionError(err)
			}
			zap.L().Error("[Audience] Get region failed", zap.Error(err))
			return serializer.DBError(err)
		}
		region = r.Country
		regionId = r.Id
	}

	var gender string
	var genderId uint
	if t.Gender != nil {
		g, err := tools.GetGender(uint(*t.Gender))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return serializer.NotFoundGenderError(err)
			}
			zap.L().Error("[Audience] Get gender failed", zap.Error(err))
			return serializer.DBError(err)
		}
		gender = g.Value
		genderId = g.Id
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
		Type:       5,
		Text:       *t.Text,
		Region:     regionId,
		Gender:     genderId,
		MinAge:     minimum,
		MaxAge:     maximum,
		LanguageId: uint(*t.Language),
	}

	tx := cache.DB.Begin()
	if err := tx.Model(&model.OptimizedAds{}).Create(&ads).Error; err != nil {
		zap.L().Error("[Audience] Create optimized ads failed", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	request.Client.Body = map[string]any{
		"text":    *t.Text,
		"region":  region,
		"gender":  gender,
		"min_age": minimum,
		"max_age": maximum,
		"lang":    language.Iso,
	}

	body, err := request.Client.Post(model.Url["generate_optimize_target_audience"], 2)
	if err != nil {
		tx.Rollback()
		return serializer.GeneratorError(err)
	}

	zap.L().Info("[Generate] Optimize Target Audience Content", zap.String("content", string(body)))

	content := model.OptimizedContent{
		Type:   5,
		AdsId:  ads.Id,
		UserId: user.Id,
		Text:   string(body),
	}
	if err := tx.Model(&model.OptimizedContent{}).Create(&content).Error; err != nil {
		zap.L().Error("[Audience] Create optimized content failed", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	tx.Commit()
	return serializer.Response{
		Code: http.StatusOK,
		Data: content,
	}
}
