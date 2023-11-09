package serializer

import "net/http"

type SummarizeHistory struct {
	Text      string `json:"text"`
	WordCount int    `json:"word_count"`
	Language  int    `json:"language"`
}

func (m *SummarizeHistory) Bind() Response {
	return Response{
		Code: http.StatusOK,
		Data: m,
	}
}
