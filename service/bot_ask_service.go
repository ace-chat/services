package service

import "ace/serializer"

type BotAsk struct {
	Content *string `form:"content" json:"content" binding:"required"`
}

func (b *BotAsk) Ask() serializer.Response {
	return serializer.Response{
		Code: 200,
		Data: "",
	}
}
