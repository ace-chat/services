package serializer

import "net/http"

type ToneHistory struct {
	Text     string `json:"text"`
	Tones    int    `json:"tones"`
	Language int    `json:"language"`
}

func (m *ToneHistory) Bind() Response {
	return Response{
		Code: http.StatusOK,
		Data: m,
	}
}
