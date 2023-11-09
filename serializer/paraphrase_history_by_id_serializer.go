package serializer

import "net/http"

type ParaphraseHistory struct {
	Text     string `json:"text"`
	Language int    `json:"language"`
}

func (m *ParaphraseHistory) Bind() Response {
	return Response{
		Code: http.StatusOK,
		Data: m,
	}
}
