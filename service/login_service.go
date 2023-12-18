package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (l *Login) Login() serializer.Response {
	var user model.User
	if err := cache.DB.Model(&model.User{}).Where("username = ?", l.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.UserNotFoundError(err)
		} else {
			zap.L().Error("[Login] Get user failed", zap.Error(err))
			return serializer.DBError(err)
		}
	}

	ok := user.CheckPassword(l.Password)
	if !ok {
		return serializer.PasswordError()
	}

	return serializer.GeneratorUser(user)
}
