package main

import (
	"fmt"
	"github.com/CareCru/interviews-go/event-hub/api"
	"net/http"
)

func main() {
	// Start Server on port 8999
	port := ":8999"
	router := api.Configure()
	fmt.Printf("Running server on port %s....", port)
	http.ListenAndServe(port, router)
}
