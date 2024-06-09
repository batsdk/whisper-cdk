package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	chi "github.com/go-chi/chi/v5"
	"net/http"
)

var chiLambda *chiadapter.ChiLambda

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	router := chi.NewRouter()

	// Define routes
	router.Get("/sample", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World from chi"))
	})

	// Use chiadapter to handle the request with chi router
	chiLambda := chiadapter.New(router)
	return chiLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(handler)
}
