package survey

// Survey is the main model which have question and answers
type Survey struct {
	ID       string   `json:"id"`
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}
