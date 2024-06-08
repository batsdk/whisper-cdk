package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"strings"
	"whisper-lambda/app"
)

func main() {
	awsApp := app.NewApp()
	lambda.Start(func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch req.Path {
		case "/sample":
			return awsApp.ApiHandler.SampleRequest(req)
		case "/groups":
			return awsApp.ApiHandler.CreateGroup(req)
		default:
			if strings.HasPrefix(req.Path, "/groups/increment") {
				return awsApp.ApiHandler.IncrementGroupMemberCount(req)
			}
			return events.APIGatewayProxyResponse{
				Body:       "Not Found",
				StatusCode: http.StatusNotFound,
			}, nil
		}
	})
}
