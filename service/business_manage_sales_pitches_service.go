package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type ManageSalesAndPitchesRequest struct {
	Id           uint                      `form:"id" json:"id" binding:"required"`
	SalesPitches []serializer.SalesPitches `form:"salesPitches" json:"salesPitches" binding:"required"`
}

func (r *ManageSalesAndPitchesRequest) Manage(user model.User) serializer.Response {
	var count int64
	if err := cache.DB.Model(&model.BusinessChatBot{}).Where("user_id = ?", user.Id).Count(&count).Error; err != nil {
		zap.L().Error("[ManageQuestionAndAnswerRequest] Get business chat bot count failed", zap.Error(err))
		return serializer.DBError(err)
	}

	if count == 0 {
		return serializer.IllegalError()
	}

	var chat model.BusinessChatBot
	if err := cache.DB.Model(&model.BusinessChatBot{}).Where("user_id = ?", user.Id).Last(&chat).Error; err != nil {
		zap.L().Error("[ManageQuestionAndAnswerRequest] Get business chat bot failed", zap.Error(err))
		return serializer.DBError(err)
	}

	sp := model.BusinessChatBotSalesPitch{
		BusinessBotId: chat.Id,
	}
	if err := cache.DB.Model(&model.BusinessChatBotSalesPitch{}).Delete(&sp).Error; err != nil {
		zap.L().Error("[ManageQuestionAndAnswerRequest] Delete business chat bot sales pitch failed", zap.Error(err))
		return serializer.DBError(err)
	}

	sps := make([]model.BusinessChatBotSalesPitch, 0)
	for _, pitch := range r.SalesPitches {
		sps = append(sps, model.BusinessChatBotSalesPitch{
			BusinessBotId: chat.Id,
			Topic:         pitch.Topic,
			Input:         pitch.Input,
		})
	}
	if err := cache.DB.Model(&model.BusinessChatBotSalesPitch{}).Create(&sps).Error; err != nil {
		zap.L().Error("[ManageQuestionAndAnswerRequest] Create business chat bot sales pitch failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
	}
}
