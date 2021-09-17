package survey

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	t.Log("this is a test")

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
	}

	for i := range samples {
		surveyRepoTMP.Create(samples[i].In)
		time.Sleep(1 * time.Second)
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
