package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"whisper-lambda/app"
)

type ReqEvent struct {
	Body string `json:"body"`
}

func sampleRequest(req ReqEvent) (string, error) {
	if req.Body == "" {
		return "", fmt.Errorf("body is empty")
	}
	return fmt.Sprint("body was = %s", req.Body), nil
}

func main() {
	awsApp := app.NewApp()
	lambda.Start(func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch req.Path {
		case "/sample":
			return awsApp.ApiHandler.SampleRequest(req)
		default:
			return events.APIGatewayProxyResponse{
				Body:       "Not Found",
				StatusCode: http.StatusNotFound,
			}, nil
		}
	})
}
