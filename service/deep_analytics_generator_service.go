package service

import (
	"ace/cache"
	"ace/model"
	"ace/request"
	"ace/serializer"
	"ace/utils"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DeepAnalytics struct {
	BusinessDescription *string `form:"business_description" json:"business_description" binding:"required"`
	ProductDescription  *string `form:"product_description" json:"product_description" binding:"required"`
	DataDescription     *string `form:"data_description" json:"data_description" binding:"required"`
	ServiceId           *uint   `form:"service_id" json:"service_id" binding:"required"`
	Filename            *string `form:"filename" json:"filename" binding:"required"`
}

func (d *DeepAnalytics) Generator(user model.User) serializer.Response {
	if *d.BusinessDescription == "" || *d.ProductDescription == "" || *d.DataDescription == "" || *d.ServiceId == 0 || *d.Filename == "" {
		zap.L().Error("[DeepAnalytics] Parameter cannot be empty")
		return serializer.ParamError(errors.New("parameter cannot be empty"))
	}

	var tool utils.Common

	services, err := tool.GetService(*d.ServiceId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.NotFoundServiceError(err)
		}
		zap.L().Error("[DeepAnalytics] Get service failed", zap.Error(err))
		return serializer.DBError(err)
	}

	analytics := model.Analytics{
		UserId:       user.Id,
		Type:         2,
		Title:        fmt.Sprintf("Deep analytics for %v", *d.Filename),
		BusinessDesc: *d.BusinessDescription,
		ProductDesc:  *d.ProductDescription,
		DataDesc:     *d.DataDescription,
		ServiceId:    services.Id,
	}

	body, err := request.Client.Post(*d.Filename, true)
	if err != nil {
		zap.L().Error("[DeepAnalytics] Create deep analytics failed", zap.Error(err))
		return serializer.GeneratorError(err)
	}
	analytics.Content = string(body)

	if err := cache.DB.Model(&model.Analytics{}).Create(&analytics).Error; err != nil {
		zap.L().Error("[DeepAnalytics] Create deep analytics record failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: 200,
		Data: analytics,
	}
}
