package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"ace/utils"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"regexp"
)

type BusinessCreateRequest struct {
	CompanyName         string                         `form:"company_name" json:"company_name" binding:"required"`
	Links               []string                       `form:"links" json:"links" binding:"required"`
	CompanyIntroduction string                         `form:"company_introduction" json:"company_introduction" binding:"required"`
	Platform            int                            `form:"platform" json:"platform" binding:"required"`
	PhoneNumber         string                         `form:"phone_number" json:"phone_number" binding:"required"`
	Tone                int                            `form:"tone" json:"tone" binding:"required"`
	QA                  []serializer.QuestionAndAnswer `form:"qa" json:"qa" binding:"required"`
	SalesPitches        []serializer.SalesPitches      `form:"sales_pitches" json:"sales_pitches"`
	Files               []string                       `form:"files" json:"files"`
}

func (r *BusinessCreateRequest) CreateBusinessBot(user model.User) serializer.Response {
	if len(r.Links) == 0 {
		return serializer.LinksError(nil)
	} else {
		pattern := "(https)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]"
		for _, link := range r.Links {
			match, err := regexp.MatchString(pattern, link)
			if err != nil {
				zap.L().Error("[BusinessCreateRequest] Match string failed", zap.Error(err))
				return serializer.LinksError(err)
			}
			if !match {
				return serializer.LinksError(nil)
			}
		}
	}

	if len(r.QA) == 0 {
		return serializer.QAError(nil)
	} else {
		for _, qa := range r.QA {
			if qa.Question == "" || qa.Answer == "" {
				return serializer.QAError(nil)
			}
		}
	}

	var tools utils.Common

	tone, err := tools.GetTone(uint(r.Tone))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundToneError(err)
		}
		zap.L().Error("[BusinessCreateRequest] Get tone failed", zap.Error(err))
		return serializer.DBError(err)
	}

	tx := cache.DB.Begin()

	var count int64
	if err := tx.Model(&model.BusinessChatBot{}).Where("user_id = ?", user.Id).Count(&count).Error; err != nil {
		zap.L().Error("[BusinessCreateRequest] Get business chat bot count failed", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	if count != 0 {
		tx.Rollback()
		return serializer.IllegalError()
	}

	businessBot := model.BusinessChatBot{
		UserId:              user.Id,
		CompanyName:         r.CompanyName,
		CompanyIntroduction: r.CompanyIntroduction,
		PhoneNumber:         r.PhoneNumber,
		ToneId:              tone.Id,
	}

	if err := tx.Model(&model.BusinessChatBot{}).Create(&businessBot).Error; err != nil {
		zap.L().Error("[BusinessCreateRequest] Create business bot failed", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	platforms := make([]model.Platform, 0)
	if err := tx.Model(&model.Platform{}).Where("type = ?", 3).Find(&platforms).Error; err != nil {
		zap.L().Error("[BusinessCreateRequest] Find platforms failed", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	p := make([]model.BusinessChatBotPlatform, 0)

	for _, platform := range platforms {
		var status bool
		if platform.Id == uint(r.Platform) {
			status = true
		}

		p = append(p, model.BusinessChatBotPlatform{
			BusinessBotId: businessBot.Id,
			Platform:      platform.Id,
			Status:        status,
		})
	}

	if err := tx.Model(&model.BusinessChatBotPlatform{}).Create(&p).Error; err != nil {
		zap.L().Error("[BusinessCreateRequest] Create business chat bot platform failed", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	links := make([]model.BusinessChatBotLink, 0)
	for _, link := range r.Links {
		links = append(links, model.BusinessChatBotLink{
			BusinessBotId: businessBot.Id,
			Url:           link,
		})
	}

	if err := tx.Model(&model.BusinessChatBotLink{}).Create(&links).Error; err != nil {
		zap.L().Error("[BusinessCreateRequest] Create link failed", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	questionAndAnswer := make([]model.BusinessChatBotQA, 0)
	for _, qa := range r.QA {
		questionAndAnswer = append(questionAndAnswer, model.BusinessChatBotQA{
			BusinessBotId: businessBot.Id,
			Question:      qa.Question,
			Answer:        qa.Answer,
		})
	}

	if err := tx.Model(&model.BusinessChatBotQA{}).Create(&questionAndAnswer).Error; err != nil {
		zap.L().Error("[BusinessCreateRequest] Create BusinessChatBotQA failed", zap.Error(err))
		tx.Rollback()
		return serializer.DBError(err)
	}

	if len(r.SalesPitches) != 0 {
		sp := make([]model.BusinessChatBotSalesPitch, 0)
		for _, pitch := range r.SalesPitches {
			sp = append(sp, model.BusinessChatBotSalesPitch{
				BusinessBotId: businessBot.Id,
				Topic:         pitch.Topic,
				Input:         pitch.Input,
			})
		}

		if err := tx.Model(&model.BusinessChatBotSalesPitch{}).Create(&sp).Error; err != nil {
			zap.L().Error("[BusinessCreateRequest] Create sales pitches failed", zap.Error(err))
			tx.Rollback()
			return serializer.DBError(err)
		}
	}

	if len(r.Files) != 0 {
		f := make([]model.BusinessChatBotFile, 0)
		for _, file := range r.Files {
			f = append(f, model.BusinessChatBotFile{
				BusinessBotId: businessBot.Id,
				Url:           file,
			})
		}

		if err := tx.Model(&model.BusinessChatBotFile{}).Create(&f).Error; err != nil {
			zap.L().Error("[BusinessCreateRequest] Create BusinessChatBotFile failed", zap.Error(err))
			tx.Rollback()
			return serializer.DBError(err)
		}
	}

	tx.Commit()

	return serializer.Response{
		Code: 200,
		Data: true,
	}
}
