# StanApi

AWS Lambda version of the Stan task. Written in Go, deployed on AWS Lambda with AWS API Gateway

https://6ly9iav2pc.execute-api.ap-southeast-2.amazonaws.com/v1

## How to build 

`GOOS=linux go build main.go`

## How to prepare AWS Lambda zip folder

`zip main.zip main`

## How to test

`go test ./...`

## How to see test coverage

`go test -cover ./...`
