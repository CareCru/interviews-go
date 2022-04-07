package api

import (
	"../service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, found := params["userID"]
	if !found {
		w.WriteHeader(http.StatusNotFound)
	}

	user, err := service.GetUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	jsonResponse, jsonError := json.Marshal(user)
	if jsonError != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
