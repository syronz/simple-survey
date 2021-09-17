package answer

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

// Create is used for creating a answer
func (p *API) Create(c *gin.Context) {
	var ans Answer

	if err := c.ShouldBindJSON(&ans); err != nil {
		log.Println("error in binding answer")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "error in bindig answer, make sure send question and answer correctly",
		})
		return
	}

	if err := p.Repo.Create(ans); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "answer successfully recorded",
	})
}
