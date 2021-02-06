package errors

import(
	"../src"
	"github.com/aws/aws-lambda-go/events"
)

func HandleBadRequestError() events.APIGatewayProxyResponse {
	body := map[string]string{"error": "Could not decode request: JSON parsing failed"}
	statusCode := 400
	resp := src.CreateResponse(body, statusCode)
	return resp
}