package service

import "github.com/CareCru/interviews-go/event-hub/repository"

func GetEventsForUser(userID string) []repository.Event {
	return repository.GetEventsForUser(userID)
}
