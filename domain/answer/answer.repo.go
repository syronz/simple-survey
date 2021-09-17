package answer

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Repo save the data to the storage
type Repo struct {
	Path string
}

// ProvideRepo return an instance of sruvey.Repo
func ProvideRepo(path string) Repo {
	return Repo{Path: path}
}

// Create is used for appending an answer to question file
func (p *Repo) Create(ans Answer) (err error) {

	var answerFile *os.File
	if answerFile, err = os.OpenFile(filepath.Join(p.Path, ans.Question),
		os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644); err != nil {
		return
	}

	defer answerFile.Close()

	now := time.Now().Unix()
	str := fmt.Sprintf("%v,%v\n", now, ans.Answer)
	if _, err = answerFile.WriteString(str); err != nil {
		return
	}

	answerFile.Sync()

	return
}
