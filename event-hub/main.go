package main

import (
	"github.com/interviews-go/event-hub/api"
	"net/http"
)

func main() {
	// Start Server on port 8999
	router := api.Configure()
	http.ListenAndServe(":8999", router)
}
