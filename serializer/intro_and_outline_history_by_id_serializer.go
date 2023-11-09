package serializer

import "net/http"

type IntroAndOutlineHistory struct {
	Topic      string `json:"topic"`
	Tones      int    `json:"tones"`
	BrandVoice int    `json:"brand_voice"`
	MinAge     int    `json:"min_age"`
	MaxAge     int    `json:"max_age"`
	Language   int    `json:"language"`
}

func (m *IntroAndOutlineHistory) Bind() Response {
	return Response{
		Code: http.StatusOK,
		Data: m,
	}
}
