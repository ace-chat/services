package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
)

type GenderRequest struct{}

func (g *GenderRequest) GetGender() serializer.Response {
	genders := make([]model.Gender, 0)
	if err := cache.DB.Model(&model.Gender{}).Find(&genders).Error; err != nil {
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: 200,
		Data: genders,
	}
}
