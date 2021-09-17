package survey

// Survey is the main model which have question and answers
type Survey struct {
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}
