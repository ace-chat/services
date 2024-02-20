package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type ManageQuestionAndAnswerRequest struct {
	Id uint                           `form:"id" json:"id" binding:"required"`
	QA []serializer.QuestionAndAnswer `form:"qa" json:"qa" binding:"required"`
}

func (r *ManageQuestionAndAnswerRequest) Manage(user model.User) serializer.Response {
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

	f := model.BusinessChatBotQA{
		BusinessBotId: chat.Id,
	}
	if err := cache.DB.Model(&model.BusinessChatBotQA{}).Where("business_bot_id = ?", chat.Id).Delete(&f).Error; err != nil {
		zap.L().Error("[ManageQuestionAndAnswerRequest] Delete business chat bot question and answer failed", zap.Error(err))
		return serializer.DBError(err)
	}

	qas := make([]model.BusinessChatBotQA, 0)
	for _, questionAndAnswer := range r.QA {
		qas = append(qas, model.BusinessChatBotQA{
			BusinessBotId: chat.Id,
			Question:      questionAndAnswer.Question,
			Answer:        questionAndAnswer.Answer,
		})
	}
	if err := cache.DB.Model(&model.BusinessChatBotQA{}).Create(&qas).Error; err != nil {
		zap.L().Error("[ManageQuestionAndAnswerRequest] Create business chat bot question and answer failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
	}
}
