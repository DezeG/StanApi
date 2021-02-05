package main

import(
	"github.com/aws/aws-lambda-go/lambda"
	"./structs"
)

func HandleLambdaEvent(req structs.StanRequest) ([]structs.StanResponse, error) {
	var resp []structs.StanResponse
	for _, movie := range req.Payload {
		if movie.Drm == true {
			var buf structs.StanResponse = structs.StanResponse{}
			buf.Image = movie.Image.ShowImage
			buf.Slug = movie.Slug
			buf.Title = movie.Title
			resp = append(resp, buf)
		}
	}
	return resp, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}