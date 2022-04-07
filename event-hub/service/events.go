package service

import (
	"github.com/CareCru/interviews-go/event-hub/repository"
	"github.com/pkg/errors"
)

func GetEventsForUser(userID string) []repository.Event {
	return repository.GetEventsForUser(userID)
}

func CreateEventForUser(userID string, input repository.Event) error {
	userEvents := repository.GetEventsForUser(userID)

	for _, event := range userEvents {
		if event.Name == input.Name {
			return errors.New("Event with same name already exists for this user")
		}
	}

	repository.CreateEvent(input)
	return nil
}
