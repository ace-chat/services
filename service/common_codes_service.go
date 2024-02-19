package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"ace/utils"
	"go.uber.org/zap"
)

type CodeRequest struct{}

func (r *CodeRequest) GetCodes() serializer.Response {
	regions := make([]model.Region, 0)
	if err := cache.DB.Model(&model.Region{}).Find(&regions).Error; err != nil {
		zap.L().Error("[GetCodes] Find region failed", zap.Error(err))
		return serializer.DBError(err)
	}

	codes := make([]string, 0)
	for _, region := range regions {
		codes = append(codes, region.Code)
	}

	return serializer.Response{
		Code: 200,
		Data: utils.Deduplicate(codes),
	}
}
