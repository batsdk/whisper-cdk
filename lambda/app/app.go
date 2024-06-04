package app

import (
	"whisper-lambda/api"
	"whisper-lambda/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	db := database.NewDynamoDBInstance()
	apiInstance := api.NewApiHandler(db)

	return App{
		ApiHandler: apiInstance,
	}
}
