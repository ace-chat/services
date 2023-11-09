package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"
)

type RegionRequest struct{}

func (r *RegionRequest) GetRegions() serializer.Response {
	regions := make([]model.Region, 0)
	if err := cache.DB.Model(&model.Region{}).Find(&regions).Error; err != nil {
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: regions,
	}
}
