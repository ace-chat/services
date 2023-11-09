package service

import (
	"ace/model"
	"ace/serializer"
	"net/http"
)

type GetUserInfo struct{}

func (g *GetUserInfo) GetUserInfo(user model.User) serializer.Response {
	return serializer.Response{
		Code: http.StatusOK,
		Data: user,
	}
}
