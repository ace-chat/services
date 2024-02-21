package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type GetChatBotRequest struct{}

func (r *GetChatBotRequest) Get(user model.User) serializer.Response {
	var count int64
	if err := cache.DB.Model(&model.BusinessChatBot{}).Where("user_id = ?", user.Id).Count(&count).Error; err != nil {
		zap.L().Error("[GetChatBotRequest] Get business chat bot count failed", zap.Error(err))
		return serializer.DBError(err)
	}

	if count == 0 {
		return serializer.Response{
			Code: 200,
			Data: true,
		}
	}

	var chat model.BusinessChatBot
	if err := cache.DB.Model(&model.BusinessChatBot{}).Where("user_id = ?", user.Id).Last(&chat).Error; err != nil {
		zap.L().Error("[GetChatBotRequest] Get business chat bot failed", zap.Error(err))
		return serializer.DBError(err)
	}

	platforms := make([]model.BusinessChatBotPlatform, 0)
	if err := cache.DB.Model(&model.BusinessChatBotPlatform{}).Where("business_bot_id = ?", chat.Id).Find(&platforms).Error; err != nil {
		zap.L().Error("[GetChatBotRequest] Find business chat bot platforms failed", zap.Error(err))
		return serializer.DBError(err)
	}

	ps := make([]uint, 0)
	for _, platform := range platforms {
		ps = append(ps, platform.Platform)
	}

	qas := make([]serializer.QuestionAndAnswer, 0)
	if err := cache.DB.Model(&model.BusinessChatBotQA{}).Where("business_bot_id = ?", chat.Id).Find(&qas).Error; err != nil {
		zap.L().Error("[GetChatBotRequest] Find business chat bot question and answer failed", zap.Error(err))
		return serializer.DBError(err)
	}

	sps := make([]serializer.SalesPitches, 0)
	if err := cache.DB.Model(&model.BusinessChatBotSalesPitch{}).Where("business_bot_id = ?", chat.Id).Find(&sps).Error; err != nil {
		zap.L().Error("[GetChatBotRequest] Find business chat bot sales and pitches failed", zap.Error(err))
		return serializer.DBError(err)
	}

	links := make([]model.BusinessChatBotLink, 0)
	if err := cache.DB.Model(&model.BusinessChatBotLink{}).Where("business_bot_id = ?", chat.Id).Find(&links).Error; err != nil {
		zap.L().Error("[GetChatBotRequest] Find business chat bot links failed", zap.Error(err))
		return serializer.DBError(err)
	}

	ls := make([]string, 0)
	for _, link := range links {
		ls = append(ls, link.Url)
	}

	files := make([]model.BusinessChatBotFile, 0)
	if err := cache.DB.Model(&model.BusinessChatBotFile{}).Where("business_bot_id = ?", chat.Id).Find(&files).Error; err != nil {
		zap.L().Error("[GetChatBotRequest] Find business chat bot files failed", zap.Error(err))
		return serializer.DBError(err)
	}

	fs := make([]string, 0)
	for _, file := range files {
		fs = append(fs, file.Url)
	}

	chatBot := serializer.BusinessChatBot{
		Id:                  chat.Id,
		CompanyName:         chat.CompanyName,
		Links:               ls,
		CompanyIntroduction: chat.CompanyIntroduction,
		Platform:            ps,
		PhoneNumber:         chat.PhoneNumber,
		Tone:                chat.ToneId,
		QA:                  qas,
		SalesPitches:        sps,
		Files:               fs,
	}

	return chatBot.Bind()
}
