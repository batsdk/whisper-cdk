package api

import (
	"encoding/json"
	"net/http"
	"whisper-lambda/types"
	"whisper-lambda/utils"

	"github.com/aws/aws-lambda-go/events"
)

type ApiHandler struct {
	dbStore types.IDatabase
}

func NewApiHandler(databaseStore types.IDatabase) ApiHandler {
	return ApiHandler{
		dbStore: databaseStore,
	}
}

func (api ApiHandler) SampleRequest(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World from chi in API Handler"))
}

//func (api ApiHandler) SampleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
//	return events.APIGatewayProxyResponse{
//		StatusCode: http.StatusOK,
//		Body:       "Sample Request Response is going okay",
//	}, nil
//}

func (api ApiHandler) CreateGroup(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var group types.Group

	err := json.Unmarshal([]byte(request.Body), &group)

	if err != nil {
		return utils.CreateResponse(http.StatusBadRequest, "Invalid Request"), err
	}

	if group.GroupName == "" || group.CreatedBy == "" || group.GroupID == "" {
		return utils.CreateResponse(http.StatusBadRequest, "Fields Cannot be empty"), err
	}

	err = api.dbStore.CreateGroup(group)

	if err != nil {
		return utils.CreateResponse(http.StatusInternalServerError, "Error creating group"), err
	}

	return utils.CreateResponse(http.StatusCreated, "Group Created"), nil

}

func (api ApiHandler) IncrementGroupMemberCount(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	groupID := request.PathParameters["id"]

	if groupID == "" {
		return utils.CreateResponse(http.StatusBadRequest, "Group ID Can not be empty"), nil
	}

	err := api.dbStore.IncrementGroupMemberCount(groupID)

	if err != nil {
		return utils.CreateResponse(http.StatusInternalServerError, "Error incrementing group member count"), err
	}

	return utils.CreateResponse(http.StatusCreated, "Group Member count updated"), nil

}
