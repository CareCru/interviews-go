package service

import (
	"github.com/CareCru/interviews-go/event-hub/repository"
)

func GetUser(id string) (repository.User, error) {
	return repository.GetUser(id)
}
func GetAllUsers() []repository.User {
	return repository.GetAllUsers()
}
