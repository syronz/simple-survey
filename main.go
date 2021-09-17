package main

import (
	"fmt"
	"simplesurvey/domain/survey"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("vim-go")

	r := gin.Default()

	rg := r.Group("/api/v1")

	surveyAPI := new(survey.SurveyAPI)
	rg.GET("/ping", surveyAPI.Ping)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
