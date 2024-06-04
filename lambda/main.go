package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
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
	lambda.Start(sampleRequest)
}
