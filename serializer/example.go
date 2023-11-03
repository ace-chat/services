package serializer

import "net/http"

type Example struct {
	Code   string `json:"code"`
	Number int    `json:"number"`
}

func (e *Example) Example(example Example) Response {
	return Response{
		Code: http.StatusOK,
		Data: example,
	}
}
