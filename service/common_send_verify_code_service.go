package service

import (
	"ace/model"
	"ace/serializer"
)

type SendVerifyCodeRequest struct {
	Type   string `form:"type" json:"type" binding:"required"`
	Target string `form:"target" json:"target"`
}

func (r *SendVerifyCodeRequest) Send(user model.User) serializer.Response {
	switch r.Type {
	case "email":
		if r.Target != "" {
			return serializer.IllegalError()
		}
		// TODO send email message
	case "phone":
		if r.Target == "" {
			return serializer.IllegalError()
		}
		// TODO send phone message
	default:
		return serializer.Response{
			Code: 200,
			Data: false,
		}
	}

	return serializer.Response{
		Code: 200,
		Data: true,
	}
}
