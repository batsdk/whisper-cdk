package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/go-chi/chi/v5"
	"whisper-lambda/app"
)

var chiLambda *chiadapter.ChiLambda

func init() {
	router := chi.NewRouter()
	api := app.NewApp()

	// Define routes
	router.Get("/sample", api.ApiHandler.SampleRequest)
	router.Post("/groups", api.ApiHandler.CreateGroup)
	router.Post("/groups/increment/{id}", api.ApiHandler.IncrementGroupMemberCount)

	// Initialize ChiLambda with Router
	chiLambda = chiadapter.New(router)
}

func handleConnect(req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Handle WebSocket connect
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "Connected"}, nil
}

func handleDisconnect(req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Handle WebSocket disconnect
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "Disconnected"}, nil
}

func handleSendMessage(req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Handle WebSocket sendMessage route
	var message map[string]interface{}
	if err := json.Unmarshal([]byte(req.Body), &message); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid message format"}, err
	}
	// Process the message
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "Message received"}, nil
}

func handleWebsocket(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.RequestContext.RouteKey {
	case "$connect":
		return handleConnect(req)
	case "$disconnect":
		return handleDisconnect(req)
	case "sendMessage":
		return handleSendMessage(req)
	default:
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid route key"}, nil
	}
}

func handler(ctx context.Context, event interface{}) (interface{}, error) {
	switch req := event.(type) {
	case events.APIGatewayProxyRequest:
		return chiLambda.ProxyWithContext(ctx, req)
	case events.APIGatewayWebsocketProxyRequest:
		return handleWebsocket(ctx, req)
	default:
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid request err in 68"}, nil
	}
}

func main() {
	lambda.Start(handler)
}
