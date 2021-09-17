package survey

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// API is used for control the API
type API struct {
	Repo Repo
}

// ProvideAPI return an instance of sruvey.API
func ProvideAPI(repo Repo) API {
	return API{Repo: repo}
}

func (p *API) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Create is used for creating a survey
func (p *API) Create(c *gin.Context) {
	var sur Survey

	if err := c.ShouldBindJSON(&sur); err != nil {
		log.Println("error in binding survey")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "error in bindig survey, make sure send question and answers correctly",
		})
		return
	}

	if err := p.Repo.Create(sur); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "survey successfully created",
	})
}
