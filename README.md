# simple-survey
simple-survey is an API that save data to file developed in DDD (Domain Driven Design), 
almost!

[![Go Report Card](https://goreportcard.com/badge/github.com/syronz/simple-survey)](https://goreportcard.com/report/github.com/syronz/simple-survey)

## Approaches
This survey can be completed in three ways (without database and saving in storage)
- implementing each question in a separate file and also answers in separate files based on the question's name
- use a csv to hold all questions as surveys with answers. save answers in different file or follow
    special approach to save questions and result in one file. One row for question and answer and
    another row for keeping the number of answers that are selected. for instance:
``` 
q1,a1,a2,a3,a4
"",15,10,19,0
q2,a1,a2
"",95,18
```
- save all state in a struct and after any request save the file serialized in gob format and each
    time app started, load that file and continue as a cache in memory

The first approach was chosen because reading from and writing to storage became minimal 
and because each question saved in a different file the lock file happened less compared 
to other solutions. The use case determines what to do, by the way. In some cases, the third 
option may be the best. We consider not using a database. With a database the approach will be different.



## Run
```
go mod tidy
go run main.go
```

## Test

For run all unit test you can use below command, **test considered in Unix environment**

```bash
go test ./...
```

After running the app you can import **survey.postman_collection.json** or run below curl
comands for testing the endpoints.

### Create survey

```bash
curl --location --request POST 'http://127.0.0.1:8080/api/v1/surveys' \
--header 'Content-Type: application/json' \
--data-raw '{
    "question": "What is your favorite color?",
    "answers": ["Blue","Green","Red", "Yellow"]
}'
```

### Answer survey
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


### Retriving the results

```bash
curl --location --request GET 'http://127.0.0.1:8080/api/v1/results' \
--header 'Content-Type: application/json' \
--data-raw '{
    "question": "What is your favorite color?",
    "answer": "Blue"
}'
```
