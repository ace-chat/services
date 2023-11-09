package serializer

import "net/http"

type EntireHistory struct {
	IntroAndOutlineHistory
	Keyword      string `json:"keyword"`
	WordCount    int    `json:"word_count"`
	OtherDetails string `json:"other_details"`
}

func (m *EntireHistory) Bind() Response {
	return Response{
		Code: http.StatusOK,
		Data: m,
	}
}
