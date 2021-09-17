# simple-survey
simple-survey is an API that save data to file developed in DDD (Domain Driven Design), 
almost!

[![Go Report Card](https://goreportcard.com/badge/github.com/syronz/simple-survey)](https://goreportcard.com/report/github.com/syronz/simple-survey)


# Run
```
go mod tidy
go run main.go
```

# Test

For run all unit test you can use below command, **test considered in Unix environment**

```bash
go test ./...
```

After running the app you can import **survey.postman_collection.json** or run below curl
comands for testing the endpoints.

## Create survey

```bash
curl --location --request POST 'http://127.0.0.1:8080/api/v1/surveys' \
--header 'Content-Type: application/json' \
--data-raw '{
    "question": "What is your favorite color?",
    "answers": ["Blue","Green","Red", "Yellow"]
}'
```

## Answer survey
for answering the survey question of the survey and survey should be send together, I didn't
implement id for surveys.

```bash
curl --location --request POST 'http://127.0.0.1:8080/api/v1/answers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "question": "What is your favorite color?",
    "answer": "Blue"
}'
```


## Retriving the results

```bash
curl --location --request GET 'http://127.0.0.1:8080/api/v1/results' \
--header 'Content-Type: application/json' \
--data-raw '{
    "question": "What is your favorite color?",
    "answer": "Blue"
}'
```
