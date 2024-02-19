package serializer

type QuestionAndAnswer struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type SalesPitches struct {
	Topic string `json:"topic"`
	Input string `json:"input"`
}
