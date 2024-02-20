package serializer

type QuestionAndAnswer struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type SalesPitches struct {
	Topic string `json:"topic"`
	Input string `json:"input"`
}

type BusinessChatBot struct {
	Id                  uint                `json:"id"`
	CompanyName         string              `json:"company_name"`
	Links               []string            `json:"links"`
	CompanyIntroduction string              `json:"company_introduction"`
	Platform            []uint              `json:"platform"`
	PhoneNumber         string              `json:"phone_number"`
	Tone                uint                `json:"tone"`
	QA                  []QuestionAndAnswer `json:"qa"`
	SalesPitches        []SalesPitches      `json:"sales_pitches"`
	Files               []string            `json:"files"`
}

func (b *BusinessChatBot) Bind() Response {
	return Response{
		Code: 200,
		Data: b,
	}
}
