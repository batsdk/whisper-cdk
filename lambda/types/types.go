package types

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type Group struct {
	GroupName string `json:"groupName"`
	GroupID   string `json:"groupID"`
	CreatedBy string `json:"createdBy"`
}

type IApiEvents interface {
	SampleRequest(w http.ResponseWriter, req *http.Request)
	CreateGroup(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	IncrementGroupMemberCount(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

type IDatabase interface {
	CreateGroup(group Group) error
	IncrementGroupMemberCount(groupID string) error
}
