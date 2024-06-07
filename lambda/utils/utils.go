package utils

import "github.com/aws/aws-lambda-go/events"

func CreateResponse(statusCode int, message string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       message,
	}
}
