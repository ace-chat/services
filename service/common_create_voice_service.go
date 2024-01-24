package service

import (
	"ace/model"
	"ace/request"
	"ace/serializer"
	"net/http"

	"go.uber.org/zap"
)

type CommonCreateVoiceRequest struct {
	Text string `form:"text" json:"text" binding:"required"`
}

func (c *CommonCreateVoiceRequest) CreateVoice(user model.User) serializer.Response {
	voice := model.Voice{
		UserId: user.Id,
		Text:   c.Text,
	}

	request.Client.Body = map[string]any{
		"brand_text": voice.Text,
	}
	body, err := request.Client.Post(model.Url["create_brand_voice"], 2)
	if err != nil {
		zap.L().Error("[Common] Create brand voice failed", zap.Error(err))
		return serializer.GeneratorError(err)
	}
	voice.Content = string(body)

	return serializer.Response{
		Code: http.StatusOK,
		Data: voice.Content,
	}
}
