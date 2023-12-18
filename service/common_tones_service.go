package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
	"net/http"
)

type TonesRequest struct {
	Type int `form:"type" json:"type" binding:"required"`
}

func (t *TonesRequest) GetTones() serializer.Response {
	tones := make([]model.Tone, 0)
	if err := cache.DB.Model(&model.Tone{}).Where("type = ?", t.Type).Find(&tones).Error; err != nil {
		zap.L().Error("[Common] Get tones failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: tones,
	}
}
