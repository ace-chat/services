package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type ChangeBusinessPhoneNumberRequest struct {
	Id         uint   `form:"id" json:"id" binding:"required"`
	Phone      string `form:"phone" json:"phone" binding:"required"`
	VerifyCode string `form:"verifyCode" json:"verifyCode" binding:"required"`
}

func (r *ChangeBusinessPhoneNumberRequest) Change(user model.User) serializer.Response {
	// TODO check verify code

	if err := cache.DB.Model(&model.BusinessChatBot{}).Where("id = ? AND user_id = ?", r.Id, user.Id).Update("phone_number", r.Phone).Error; err != nil {
		zap.L().Error("[ChangeBusinessPhoneNumberRequest] Update business chat bot phone number failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
	}
}
