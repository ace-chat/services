package service

import (
	"ace/cache"
	"ace/model"
	"ace/serializer"
	"errors"
	"gorm.io/gorm"
)

type Login struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	GoogleKey string `form:"google_key" json:"google_key" binding:"required"`
}

func (l *Login) Login() serializer.Response {
	var user model.User
	if err := cache.DB.Model(&model.User{}).Where("username = ?", l.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.UserNotFoundError(err)
		} else {
			return serializer.DBError(err)
		}
	}

	ok := user.CheckPassword(l.Password)
	if !ok {
		return serializer.PasswordError()
	}

	return serializer.GeneratorUser(user)
}
