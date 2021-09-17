package answer

import (
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Result returns the list of questions
func (p *Repo) Result() (results []Result, err error) {

	var files []fs.FileInfo
	if files, err = ioutil.ReadDir(p.Path); err != nil {
		return
	}

	for _, file := range files {

		if !file.IsDir() && !strings.HasPrefix(file.Name(), ".") {
			f, e := os.Open(filepath.Join(p.Path, file.Name()))
			if e != nil {
				err = e
				return
			}
			defer f.Close()

			b := new(strings.Builder)
			io.Copy(b, f)
			data := b.String()
			result := p.parse(file.Name(), data)
			results = append(results, result)
		}
	}

	return
}

func (p *Repo) parse(question, data string) (result Result) {
	result.Question = question
	result.Answers = make(map[string]int)
	lines := strings.Split(data, "\n")
	for _, v := range lines {
		parts := strings.Split(v, ",")
		if len(parts) > 1 {
			if _, ok := result.Answers[parts[1]]; ok {
				result.Answers[parts[1]]++
			} else {
				result.Answers[parts[1]] = 1
			}
		}
	}

	return
}
