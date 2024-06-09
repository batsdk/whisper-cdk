package main

import (
	"context"
	"whisper-lambda/app"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	chi "github.com/go-chi/chi/v5"
)

var chiLambda *chiadapter.ChiLambda

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	router := chi.NewRouter()
	api := app.NewApp()

	// Define routes
	router.Get("/sample", api.ApiHandler.SampleRequest)
	router.Post("/groups", api.ApiHandler.CreateGroup)
	router.Post("/groups/increment/{id}", api.ApiHandler.IncrementGroupMemberCount)

	// Use chiadapter to handle the request with chi router
	chiLambda := chiadapter.New(router)
	return chiLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
