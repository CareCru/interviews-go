package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func Configure() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users/{userID}", getUser).Methods("GET")
	router.HandleFunc("/users", getAllUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{userID}/events", getEventsForUser).Methods("GET")
	router.HandleFunc("/users/{userID}/events", createEventForUser).Methods("POST")
	return router
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
