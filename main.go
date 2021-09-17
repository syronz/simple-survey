package main

import (
	"simplesurvey/domain/answer"
	"simplesurvey/domain/survey"

	"github.com/gin-gonic/gin"
)

func main() {

	// Default server runs at 8080
	r := gin.Default()

	// route group to separate versions
	rg := r.Group("/api/v1")

	// survey domain
	surveyRepo := survey.ProvideRepo("public/surveys")
	surveyAPI := survey.ProvideAPI(surveyRepo)

	// answer domain
	answerRepo := answer.ProvideRepo("public/answers")
	answerAPI := answer.ProvideAPI(answerRepo)

	// routes
	rg.POST("/surveys", surveyAPI.Create)
	rg.POST("/answers", answerAPI.Create)
	rg.GET("/results", answerAPI.Result)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
