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

	var file *os.File
	for _, v := range files {
		// if v is directory or hidden file ignore them
		if !v.IsDir() && !strings.HasPrefix(v.Name(), ".") {
			if file, err = os.Open(filepath.Join(p.Path, v.Name())); err != nil {
				return
			}
			defer file.Close()

			fileStr := new(strings.Builder)
			io.Copy(fileStr, file)
			result := parse(v.Name(), fileStr.String())
			results = append(results, result)
		}
	}

	return
}

func parse(question, data string) (result Result) {
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
