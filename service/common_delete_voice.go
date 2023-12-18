package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type CommonDeleteVoice struct {
	Id *uint `form:"id" json:"id" binding:"required"`
}

func (c *CommonDeleteVoice) Delete(user model.User) serializer.Response {
	if err := cache.DB.Model(&model.Voice{}).Where("id = ? AND user_id = ?", *c.Id, user.Id).Delete(&model.Voice{}).Error; err != nil {
		zap.L().Error("[Voice] Delete brand voice failed", zap.Error(err))
		return serializer.DBError(err)
	}
	return serializer.Response{
		Code: 200,
		Data: true,
	}
}
