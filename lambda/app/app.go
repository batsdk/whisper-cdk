package app

import (
	"whisper-lambda/api"
	"whisper-lambda/database"
	"whisper-lambda/types"
)

type App struct {
	ApiHandler types.IApiEvents
}

func NewApp() App {
	db := database.NewDynamoDBInstance()
	apiInstance := api.NewApiHandler(db)

	return App{
		ApiHandler: apiInstance,
	}
}
