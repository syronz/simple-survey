package answer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Result return the statics of each answer for questions
func (p *API) Result(c *gin.Context) {

	var results []Result
	var err error
	if results, err = p.Repo.Result(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, results)
}
