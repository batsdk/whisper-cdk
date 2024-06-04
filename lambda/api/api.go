package api

import "whisper-lambda/database"

type ApiHandler struct {
	dbStore database.IDatabase
}

func NewApiHandler(databaseStore database.IDatabase) ApiHandler {
	return ApiHandler{
		dbStore: databaseStore,
	}
}
