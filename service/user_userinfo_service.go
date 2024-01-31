package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"net/http"

	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

type GetUserInfo struct{}

type UpdateUserInfo struct {
	Username    string `form:"username json:"username"`
	DisplayName string `form:"display_name" json:"display_name"`
	Email       string `form:"email" json:"email" binding:"required"`
	Phone       string `form:"phone" json:"phone" binding:"required"`
}

type UpdatePassword struct {
	Password string `form:"password" json:"password" binding:"required"`
}

func (g *GetUserInfo) GetUserInfo(user model.User) serializer.Response {
	return serializer.Response{
		Code: http.StatusOK,
		Data: user,
	}
}

func (u *UpdateUserInfo) UpdateUserInfo(user model.User) serializer.Response {
	userModel := model.User{
		Username:    u.Username,
		DisplayName: u.DisplayName,
		Email:       u.Email,
		Phone:       u.Phone,
	}

	var updatedUser model.User
	resp := cache.DB.Model(&model.User{}).Clauses(clause.Returning{}).Where("id = ?", user.Id).Updates(&userModel).Scan(&updatedUser)
	if resp.Error != nil {
		zap.L().Error("[UpdateUserInfo] Update user failed", zap.Error(resp.Error))
		return serializer.DBError(resp.Error)
	}

	return serializer.Response{
		Code: http.StatusOK,
		Data: updatedUser,
	}
}

func (u *UpdatePassword) UpdatePassword(user model.User) serializer.Response {
	user = model.User{
		Password: u.Password,
	}

	if err := user.SetPassword(u.Password); err != nil {
		zap.L().Error("[UpdatePassword] Update password failed", zap.Error(err))
		return serializer.PasswordError()
	}

	if err := cache.DB.Model(&model.User{}).Where("id = ?", user.Id).Updates(&user).Error; err != nil {
		zap.L().Error("[UpdatePassword] Update password failed", zap.Error(err))
		return serializer.DBError(err)
	}
	return serializer.Response{
		Code: http.StatusOK,
		Data: user,
	}
}
