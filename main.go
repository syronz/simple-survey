package main

import (
	"fmt"
	"simplesurvey/domain/answer"
	"simplesurvey/domain/survey"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("vim-go")

	r := gin.Default()

	rg := r.Group("/api/v1")

	surveyRepo := survey.ProvideRepo("public/surveys")
	surveyAPI := survey.ProvideAPI(surveyRepo)

	answerRepo := answer.ProvideRepo("public/answers")
	answerAPI := answer.ProvideAPI(answerRepo)

	rg.POST("/surveys", surveyAPI.Create)
	rg.POST("/answers", answerAPI.Create)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
