package main

import(
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"./src"
	"./errors"
)

func HandleLambdaEvent(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	movies, err := src.ParseReqBody([]byte(req.Body))
	if err != nil {
		resp := errors.HandleBadRequestError()
		return resp, nil
	}
	selectedMovies := src.SelectDrmMovies(movies)
	statusCode := 200
	resp := src.CreateResponse(selectedMovies, statusCode)

	return resp, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}