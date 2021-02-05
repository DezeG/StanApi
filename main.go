package main

import(
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"./src"
)

func HandleLambdaEvent(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	movies, err := src.ParseReqBody([]byte(req.Body))
	if err != nil {
		body := map[string]string{"error": "Could not decode request: JSON parsing failed"}
		statusCode := 400
		resp := src.CreateResponse(body, statusCode)
		return resp, nil
	}
	selectedMovies := src.SelectDrmMovies(movies)
	statusCode := 200
	resp := src.CreateResponse(selectedMovies, statusCode)

	return resp, nil
/*
	return events.APIGatewayProxyResponse{
		    Headers:    map[string]string{"content-type": "application/json"},
		    Body:       string(r),
		    StatusCode: 200,} , nil
		    */

}

func main() {
	lambda.Start(HandleLambdaEvent)
}