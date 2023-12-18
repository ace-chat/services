package model

import "time"

type ChatHistory struct {
	History   string    `json:"History"`
	SessionId time.Time `json:"SessionId"`
}

type ChatSubHistory struct {
	Type string `json:"type"`
	Data struct {
		Content          string `json:"content"`
		AdditionalKwargs any    `json:"additional_kwargs"`
		Type             string `json:"type"`
		Example          bool   `json:"example"`
	} `json:"data"`
}

type History struct {
	Type    string    `json:"type"`
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}
