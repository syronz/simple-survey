package answer

import (
	"testing"
)

func TestParse(t *testing.T) {

	map1 := make(map[string]int)
	map1["A"] = 4

	map2 := make(map[string]int)
	map2["A"] = 3
	map2["B"] = 1

	samples := []struct {
		question string
		data     string
		out      Result
	}{
		{
			question: "Q1",
			data:     "1,A\n2,A\n3,A\n4,A\n",
			out: Result{
				Question: "Q1",
				Answers:  map1,
			},
		},
		{
			question: "Q2",
			data:     "1,A\n2,A\n3,A\n4,B\n",
			out: Result{
				Question: "Q2",
				Answers:  map2,
			},
		},
		{
			question: "Q3",
			data:     "",
			out: Result{
				Question: "Q3",
				Answers:  make(map[string]int),
			},
		},
	}

	for i := range samples {
		r := parse(samples[i].question, samples[i].data)
		if !compareResult(r, samples[i].out) {
			t.Errorf("for question %v, result is %v it suppose to be %v", samples[i].question,
				r, samples[i].out)
		}
	}

}

// because Result has map we need special function for compare it
func compareResult(r1, r2 Result) bool {

	if r1.Question != r2.Question {
		return false
	}

	if len(r1.Answers) != len(r2.Answers) {
		return false
	}

	for i := range r1.Answers {
		if r1.Answers[i] != r2.Answers[i] {
			return false
		}
	}

	return true
}
