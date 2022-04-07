package api

import (
	"github.com/gorilla/mux"
)

func Configure() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users/{userID}", getUser).Methods("GET")
	//router.Handle("/users", createUser).Methods("POST")
	//router.Handle("/users/{userID}/events", getEventsForUsers).Methods("GET")
	return router
}
