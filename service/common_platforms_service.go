package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type PlatformRequest struct {
	Type int `form:"type" json:"type" binding:"required"`
}

func (p *PlatformRequest) GetPlatforms() serializer.Response {
	platforms := make([]model.Platform, 0)
	if err := cache.DB.Model(&model.Platform{}).Where("type = ?", p.Type).Find(&platforms).Error; err != nil {
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: platforms,
	}
}
