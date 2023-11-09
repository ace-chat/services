package serializer

import "net/http"

type VoiceHistory struct {
	Text       string `json:"text"`
	BrandVoice int    `json:"brand_voice"`
	Language   int    `json:"language"`
}

func (m *VoiceHistory) Bind() Response {
	return Response{
		Code: http.StatusOK,
		Data: m,
	}
}
