package utils

import (
	"ace/cache"
	"ace/model"
)

type Common struct{}

func (*Common) GetPlatform(id uint) (*model.Platform, error) {
	var platform model.Platform
	if err := cache.DB.Model(&model.Platform{}).Where("id = ?", id).First(&platform).Error; err != nil {
		return nil, err
	}

	return &platform, nil
}

func (*Common) GetTone(id uint) (*model.Tone, error) {
	var tone model.Tone
	if err := cache.DB.Model(&model.Tone{}).Where("id = ?", id).First(&tone).Error; err != nil {
		return nil, err
	}

	return &tone, nil
}

func (*Common) GetVoice(id, userId uint) (*model.Voice, error) {
	var voice model.Voice
	if err := cache.DB.Model(&model.Voice{}).Where("id = ? AND user_id = ?", id, userId).First(&voice).Error; err != nil {
		return nil, err
	}
	return &voice, nil
}

func (*Common) GetRegion(id uint) (*model.Region, error) {
	var region model.Region
	if err := cache.DB.Model(&model.Region{}).Where("id = ?", id).First(&region).Error; err != nil {
		return nil, err
	}

	return &region, nil
}

func (*Common) GetGender(id uint) (*model.Gender, error) {
	var gender model.Gender
	if err := cache.DB.Model(&model.Gender{}).Where("id = ?", id).First(&gender).Error; err != nil {
		return nil, err
	}
	return &gender, nil
}

func (*Common) GetLanguage(id uint) (*model.Language, error) {
	var language model.Language
	if err := cache.DB.Model(&model.Language{}).Where("id = ?", id).First(&language).Error; err != nil {
		return nil, err
	}

	return &language, nil
}
