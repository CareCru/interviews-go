package service

import "../repository"

func GetUser(id string) (repository.User, error) {
	return repository.GetUser(id)
}
