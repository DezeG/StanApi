package src

import(
	"github.com/aws/aws-lambda-go/events"
	"log"
	"encoding/json"
)

func CreateResponse(body interface{}, statusCode int) events.APIGatewayProxyResponse {
	var resp events.APIGatewayProxyResponse
	createRespHeader(&resp)
	createRespBody(&resp, body)
	resp.StatusCode = statusCode

	return resp

}

func createRespHeader(resp *events.APIGatewayProxyResponse) {
	//resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
	resp.Headers = make(map[string]string)
	resp.Headers["Access-Control-Allow-Origin"] = "*"
	resp.Headers["Access-Control-Allow-Headers"] = "Content-Type"
	resp.Headers["Content-Type"] = "application/json"
}

func createRespBody(resp *events.APIGatewayProxyResponse, body interface{}) {
	buf, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
	}
	resp.Body = string(buf)
}