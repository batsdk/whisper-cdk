package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"whisper-lambda/types"
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

func (api ApiHandler) CreateGroup(w http.ResponseWriter, req *http.Request) {

	var group types.Group

	err := json.NewDecoder(req.Body).Decode(&group)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Can not decode body")
		return
	}

	if group.GroupName == "" || group.CreatedBy == "" || group.GroupID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Fields Cannot be empty")
		return
	}

	err = api.dbStore.CreateGroup(group)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating group")
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Group Created")

}

func (api ApiHandler) IncrementGroupMemberCount(w http.ResponseWriter, req *http.Request) {

	groupID := chi.URLParam(req, "id")

	if groupID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Fields Cannot be empty")
		return
	}

	err := api.dbStore.IncrementGroupMemberCount(groupID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating group")
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Number of group members increased")

}
