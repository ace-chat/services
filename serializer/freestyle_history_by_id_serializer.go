package serializer

import "net/http"

type FreestyleHistory struct {
	Detail     string `json:"detail"`
	Tones      int    `json:"tones"`
	BrandVoice int    `json:"brand_voice"`
	Region     int    `json:"region"`
	Gender     int    `json:"gender"`
	MinAge     int    `json:"min_age"`
	MaxAge     int    `json:"max_age"`
	Language   int    `json:"language"`
}

func (m *FreestyleHistory) Bind() Response {
	return Response{
		Code: http.StatusOK,
		Data: m,
	}
}
