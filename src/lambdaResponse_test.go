package src

import(
	"testing"
	"github.com/aws/aws-lambda-go/events"
	"../structs"
)

func TestCreateRespBody(t *testing.T) {
	var resp events.APIGatewayProxyResponse = events.APIGatewayProxyResponse{}
	testBody := structs.ResponseMovie{Image: "TestPath", Slug:"TestSlug", Title: "TestTitle"}
	CreateRespBody(&resp, testBody)
	correctBody := `{"image":"TestPath","slug":"TestSlug","title":"TestTitle"}`
	if resp.Body != correctBody {
		t.Error("Incorrect response body. Expect\n", correctBody, "\nHave\n", resp.Body)
	}
}

func  TestCreateRespHeader(t *testing.T) {
	var resp events.APIGatewayProxyResponse = events.APIGatewayProxyResponse{}
	CreateRespHeader(&resp)
	if resp.Headers["Access-Control-Allow-Origin"] != "*" ||
			resp.Headers["Access-Control-Allow-Headers"] != "Content-Type" ||
			resp.Headers["Content-Type"] != "application/json" {
		t.Error("Missing or incorrect headers")
	}
}

func TestCreateResponse(t *testing.T) {
	testBody := structs.ResponseMovie{Image: "TestPath", Slug:"TestSlug", Title: "TestTitle"}
	correctBody := `{"image":"TestPath","slug":"TestSlug","title":"TestTitle"}`
	statusCode := 123456
	resp := CreateResponse(testBody, statusCode)
	if resp.Body != correctBody {
		t.Error("Incorrect response body. Expect\n", correctBody, "\nHave\n", resp.Body)
	}
	if resp.StatusCode != statusCode {
		t.Error("Incorrect response status code. Expect\n", statusCode, "\nHave\n", resp.StatusCode)
	}
	if resp.Headers["Access-Control-Allow-Origin"] != "*" ||
			resp.Headers["Access-Control-Allow-Headers"] != "Content-Type" ||
			resp.Headers["Content-Type"] != "application/json" {
		t.Error("Missing or incorrect headers")
	}
}