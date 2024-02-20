package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"go.uber.org/zap"
)

type ManageUploadFilesRequest struct {
	Id   uint     `form:"id" json:"id" binding:"required"`
	Urls []string `form:"urls" json:"urls" binding:"required"`
}

func (r *ManageUploadFilesRequest) Manage(user model.User) serializer.Response {
	var count int64
	if err := cache.DB.Model(&model.BusinessChatBot{}).Where("user_id = ?", user.Id).Count(&count).Error; err != nil {
		zap.L().Error("[AddMoreFileRequest] Get business chat bot count failed", zap.Error(err))
		return serializer.DBError(err)
	}

	if count == 0 {
		return serializer.IllegalError()
	}

	var chat model.BusinessChatBot
	if err := cache.DB.Model(&model.BusinessChatBot{}).Where("user_id = ?", user.Id).Last(&chat).Error; err != nil {
		zap.L().Error("[AddMoreFileRequest] Get business chat bot failed", zap.Error(err))
		return serializer.DBError(err)
	}

	f := model.BusinessChatBotFile{
		BusinessBotId: chat.Id,
	}

	// delete old files
	if err := cache.DB.Model(&model.BusinessChatBotFile{}).Where("business_bot_id = ?", chat.Id).Delete(&f).Error; err != nil {
		zap.L().Error("[AddMoreFileRequest] Delete business chat bot file failed", zap.Error(err))
		return serializer.DBError(err)
	}

	// add new files
	fs := make([]model.BusinessChatBotFile, 0)
	for _, url := range r.Urls {
		fs = append(fs, model.BusinessChatBotFile{
			BusinessBotId: chat.Id,
			Url:           url,
		})
	}

	if err := cache.DB.Model(&model.BusinessChatBotFile{}).Create(&fs).Error; err != nil {
		zap.L().Error("[AddMoreFileRequest] Create business chat bot file failed", zap.Error(err))
		return serializer.DBError(err)
	}

	return serializer.Response{
		Code: 200,
		Data: true,
	}
}
