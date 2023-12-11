package serializer

import "net/http"

type SimpleAnalytics struct{}

func (s *SimpleAnalytics) Bind(data string) Response {
	return Response{
		Code: http.StatusOK,
		Data: data,
	}
}
