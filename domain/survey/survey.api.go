package survey

import "github.com/gin-gonic/gin"

type SurveyAPI struct {
}

func (p *SurveyAPI) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (p *SurveyAPI) Create(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
