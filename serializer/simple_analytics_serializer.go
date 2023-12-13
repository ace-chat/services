package serializer

import (
	"net/http"
	"time"
)

type SimpleAnalytics struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *SimpleAnalytics) Bind() Response {
	return Response{
		Code: http.StatusOK,
		Data: *s,
	}
}
