package survey

import (
	"os"
	"path/filepath"
)

// Repo save the data to the storage
type Repo struct {
	Path string
}

// ProvideRepo return an instance of sruvey.Repo
func ProvideRepo(path string) Repo {
	return Repo{Path: path}
}

// Create is used for creating a survey
func (p *Repo) Create(sur Survey) (err error) {

	var surveyFile *os.File
	if surveyFile, err = os.Create(filepath.Join(p.Path, sur.Question)); err != nil {
		return
	}

	defer surveyFile.Close()

	if _, err = surveyFile.WriteString(sur.Question + "\n"); err != nil {
		return
	}

	for _, v := range sur.Answers {
		if _, err = surveyFile.WriteString(v + "\n"); err != nil {
			return
		}
	}

	surveyFile.Sync()

	return
}
