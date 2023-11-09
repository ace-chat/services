package serializer

import "net/http"

type MediaAndEngineHistory struct {
	Platform    int    `json:"platform"`
	BrandName   string `json:"brand_name"`
	ServiceName string `json:"service_name"`
	ServiceDesc string `json:"service_desc"`
	Tones       int    `json:"tones"`
	BrandVoice  int    `json:"brand_voice"`
	Region      int    `json:"region"`
	Gender      int    `json:"gender"`
	MinAge      int    `json:"min_age"`
	MaxAge      int    `json:"max_age"`
	Language    int    `json:"language"`
	Content     string `json:"content"`
}

func (m *MediaAndEngineHistory) Bind() Response {
	return Response{
		Code: http.StatusOK,
		Data: m,
	}
}
