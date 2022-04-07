package repository

import (
	"github.com/pkg/errors"
)

var userDatabase = make(map[string]User)
var eventDatabase = make(map[string]Event)

func init() {
	userDatabase["0"] = User{
		"0",
		"Mike",
		"Fan",
		"mike@hotmail.com",
	}
	userDatabase["1"] = User{
		"1",
		"Jennifer",
		"Rockband",
		"jen@hotmail.com",
	}
	userDatabase["2"] = User{
		"2",
		"Bill",
		"Lovesmusic",
		"bill@hotmail.com",
	}

	eventDatabase["0"] = Event{
		ID:      "0",
		Name:    "Flying Cars",
		OwnerID: "0",
		Venue:   "Stanley Park",
	}

	eventDatabase["1"] = Event{
		ID:      "2",
		Name:    "Metallica Live",
		OwnerID: "0",
		Venue:   "Rodgers Arena",
	}
	eventDatabase["2"] = Event{
		ID:      "3",
		Name:    "Shakespeare",
		OwnerID: "1",
		Venue:   "Orpheum Theatre",
	}
}

// GetUser Returns user for Given ID if found
func GetUser(id string) (User, error) {
	user, foundUser := userDatabase[id]
	if !foundUser {
		return User{}, errors.New("No user found for given ID")
	}
	return user, nil
}

func GetAllUsers() []User {
	result := make([]User, len(userDatabase))
	for _, value := range userDatabase {
		result = append(result, value)
	}
	return result
}

// GetEventsForUser returns a list of events that have an ownerID that match the given user id
func GetEventsForUser(userID string) []Event {
	result := make([]Event, 0)
	for _, event := range eventDatabase {
		if event.OwnerID == userID {
			result = append(result, event)
		}
	}
	return result
}

func CreateUser(user User) {
	user.ID = string(len(userDatabase))
	userDatabase[user.ID] = user
}

func CreateEvent(event Event) {
	event.ID = string(len(eventDatabase))
	eventDatabase[event.ID] = event
}
