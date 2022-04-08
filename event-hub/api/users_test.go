package api

import (
	"bytes"
	"encoding/json"
	"github.com/CareCru/interviews-go/event-hub/api"
	"github.com/CareCru/interviews-go/event-hub/repository"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Create Event for User", func(t *testing.T) {

		payload := repository.Event{
			Name:    "API TEST EVENT",
			OwnerID: "0",
			Venue:   "Vancouver Arena",
		}

		jsonPayload, err := json.Marshal(payload)
		assert.Nil(t, err)

		req, _ := http.NewRequest("POST", "/users/2/events", bytes.NewBuffer(jsonPayload))
		response := executeRequest(req)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()

	router := api.Configure()
	router.ServeHTTP(responseRecorder, req)

	return responseRecorder
}
