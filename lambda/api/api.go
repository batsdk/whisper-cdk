package api

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"whisper-lambda/database"
)

type ApiHandler struct {
	dbStore database.IDatabase
}

func NewApiHandler(databaseStore database.IDatabase) ApiHandler {
	return ApiHandler{
		dbStore: databaseStore,
	}
}

func (api ApiHandler) SampleRequest(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Sample Request Response is going okay",
	}, nil
}
