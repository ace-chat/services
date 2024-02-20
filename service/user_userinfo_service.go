package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"

	"go.uber.org/zap"
)

type GetUserInfo struct{}

type UpdateUserInfo struct {
	DisplayName string `form:"display_name" json:"display_name"`
	Phone       string `form:"phone" json:"phone"`
}

type UpdatePassword struct {
	OldPassword string `form:"old_password" json:"old_password" binding:"required"`
	NewPassword string `form:"new_password" json:"new_password" binding:"required"`
}

func (g *GetUserInfo) GetUserInfo(user model.User) serializer.Response {
	return serializer.Response{
		Code: http.StatusOK,
		Data: user,
	}
}

func (u *UpdateUserInfo) UpdateUserInfo(user model.User) serializer.Response {
	userModel := model.User{
		DisplayName: u.DisplayName,
		Phone:       u.Phone,
	}

	resp := cache.DB.Model(&model.User{}).Where("id = ?", user.Id).Updates(&userModel)
	if resp.Error != nil {
		zap.L().Error("[UpdateUserInfo] Update user failed", zap.Error(resp.Error))
		return serializer.DBError(resp.Error)
	}

	updatedUser := model.User{}
	result := cache.DB.Where("id = ?", user.Id).First(&updatedUser)
	if result.Error != nil {
		zap.L().Error("[UpdateUserInfo] Update user failed", zap.Error(resp.Error))
		return serializer.DBError(resp.Error)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: updatedUser,
	}
}

func (u *UpdatePassword) UpdatePassword(user model.User) serializer.Response {
	if !user.CheckPassword(u.OldPassword) {
		zap.L().Error("[UpdatePassword] Old password is wrong")
		return serializer.PasswordError()
	}

	if err := user.SetPassword(u.NewPassword); err != nil {
		zap.L().Error("[UpdatePassword] Update password failed", zap.Error(err))
		return serializer.PasswordError()
	}

	userModel := model.User{
		Password: user.Password,
	}
	if err := cache.DB.Model(&model.User{}).Where("id = ?", user.Id).Updates(&userModel).Error; err != nil {
		zap.L().Error("[UpdatePassword] Update password failed", zap.Error(err))
		return serializer.DBError(err)
	}
	return serializer.Response{
		Code: http.StatusOK,
		Data: user,
	}
}
