package api

import (
	"github.com/CareCru/interviews-go/event-hub/repository"
	"github.com/gorilla/mux"
	"net/http"
)

func getEventsForUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, found := params["userID"]
	if !found {
		respondWithError(w, http.StatusBadRequest, "No userID found in request path")
		return
	}

	events := repository.GetEventsForUser(userID)
	respondWithJSON(w, http.StatusOK, events)
}
