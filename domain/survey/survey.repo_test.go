package survey

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

/*
Notice: test should be run in linux systems
*/

func TestCreate(t *testing.T) {
	const TMP = "/tmp"
	surveyRepoTMP := ProvideRepo(TMP)

	samples := []struct {
		Dir string
		In  Survey
		Out string
	}{
		{
			Dir: TMP,
			In: Survey{
				Question: "test-question",
				Answers:  []string{"A", "B", "C", "D"},
			},
			Out: "test-question\nA\nB\nC\nD\n",
		},
		{
			Dir: TMP,
			In: Survey{
				Question: "zzzzzz;;;;;;;:w",
				Answers:  []string{"A", "B", "C", "D"},
			},
			Out: "zzzzzz;;;;;;;:w\nA\nB\nC\nD\n",
		},
		{
			Dir: TMP,
			In: Survey{
				Question: "1",
				Answers:  []string{},
			},
			Out: "1\n",
		},
	}

	for i := range samples {
		surveyRepoTMP.Create(samples[i].In)
		time.Sleep(100 * time.Millisecond)
		out, _ := openFile(filepath.Join(samples[i].Dir, samples[i].In.Question))
		if out != samples[i].Out {
			t.Errorf("for test %v, the output is not same, check file %v/%v", samples[i].In.Question,
				samples[i].Dir, samples[i].In.Question)
		}
	}

}

func openFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b := new(strings.Builder)
	io.Copy(b, f)
	data := b.String()
	return data, nil
}
