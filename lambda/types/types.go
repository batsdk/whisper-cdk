package types

import (
	"net/http"
)

type Group struct {
	GroupName string `json:"groupName"`
	GroupID   string `json:"groupID"`
	CreatedBy string `json:"createdBy"`
}

type IApiEvents interface {
	SampleRequest(w http.ResponseWriter, req *http.Request)
	CreateGroup(w http.ResponseWriter, req *http.Request)
	IncrementGroupMemberCount(w http.ResponseWriter, req *http.Request)
}

type IDatabase interface {
	CreateGroup(group Group) error
	IncrementGroupMemberCount(groupID string) error
}
