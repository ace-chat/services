package serializer

import "net/http"

type MarketingAndWelcomeAndAdvantageHistory struct {
	BrandName   string `json:"brand_name"`
	ServiceName string `json:"service_name"`
	BrandDesc   string `json:"brand_desc"`
	Tones       int    `json:"tones"`
	BrandVoice  int    `json:"brand_voice"`
	Region      int    `json:"region"`
	Gender      int    `json:"gender"`
	MinAge      int    `json:"min_age"`
	MaxAge      int    `json:"max_age"`
	Language    int    `json:"language"`
}

func (m *MarketingAndWelcomeAndAdvantageHistory) Bind() Response {
	return Response{
		Code: http.StatusOK,
		Data: m,
	}
}
