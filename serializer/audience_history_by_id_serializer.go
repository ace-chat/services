package serializer

import "net/http"

type AudienceHistory struct {
	Text     string `json:"text"`
	Region   int    `json:"region"`
	Gender   int    `json:"gender"`
	MinAge   int    `json:"min_age"`
	MaxAge   int    `json:"max_age"`
	Language int    `json:"language"`
}

func (m *AudienceHistory) Bind() Response {
	return Response{
		Code: http.StatusOK,
		Data: m,
	}
}
