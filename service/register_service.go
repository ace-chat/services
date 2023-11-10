package service

import (
	"ace/model"
	"ace/serializer"
)

type RegisterRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (r *RegisterRequest) Register() serializer.Response {
	user := model.User{
		Email:       r.Email,
		Username:    r.Email,
		DisplayName: r.Email,
	}
	if err := user.SetPassword(r.Password); err != nil {
		return serializer.PasswordError()
	}

	return serializer.GeneratorUser(user)
}
