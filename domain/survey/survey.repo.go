package survey

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
func (p *Repo) Create(c *gin.Context) {
	var sur Survey

	if err := c.ShouldBindJSON(&sur); err != nil {
		log.Println("error in binding survey")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "error in bindig survey, make sure send question and answers correctly",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "survey successfully created",
	})
}
