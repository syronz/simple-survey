package answer

// Result is used for recording the answering
type Result struct {
	Question string         `json:"question"`
	Answers  map[string]int `json:"asnwers"`
}
