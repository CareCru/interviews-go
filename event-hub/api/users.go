package api

import (
	"encoding/json"
	"github.com/CareCru/interviews-go/event-hub/repository"
	"github.com/CareCru/interviews-go/event-hub/service"
	"github.com/gorilla/mux"
	"net/http"
)

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := service.GetAllUsers()
	respondWithJSON(w, http.StatusOK, users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, found := params["userID"]
	if !found {
		respondWithError(w, http.StatusNotFound, "User ID not found in params")
		return
	}

	user, err := service.GetUser(userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User Not Found")
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user repository.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	repository.CreateUser(user)
}

func createEventForUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := params["userID"]
	var event repository.Event
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&event); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	service.CreateEventForUser(userID, event)
}