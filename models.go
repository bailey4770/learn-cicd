package main

import (
	"time"

	"github.com/bootdotdev/learn-cicd-starter/internal/database"
)

/*Gosec marks exposed APIKey struct as a concern. This is valid, since original code was sharing users API key unnecessarily.
* Response only ever needs to include API key from CreateUser endpoint. Other endpoints returning user details do not need to.
* middlewareAuth func ensured user only ever gets their own api key in response, so concern is not HUGE.
* However, unnecessary key exposure only increases risk of accidental key leakage.
* Fix: create user endpoint has separate struct with exported api key. Not exported normally.*/

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	apiKey    string
}

func databaseUserToUser(user database.User) (User, error) {
	createdAt, err := time.Parse(time.RFC3339, user.CreatedAt)
	if err != nil {
		return User{}, err
	}

	updatedAt, err := time.Parse(time.RFC3339, user.UpdatedAt)
	if err != nil {
		return User{}, err
	}
	return User{
		ID:        user.ID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Name:      user.Name,
	}, nil
}

type NewUser struct {
	User
	APIKey string `json:"api_key"` // #nosec G117
}

func databaseUserToNewUser(user database.User) (NewUser, error) {
	u, err := databaseUserToUser(user)
	if err != nil {
		return NewUser{}, err
	}

	return NewUser{
		User:   u,
		APIKey: u.apiKey,
	}, nil
}

type Note struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Note      string    `json:"note"`
	UserID    string    `json:"user_id"`
}

func databaseNoteToNote(post database.Note) (Note, error) {
	createdAt, err := time.Parse(time.RFC3339, post.CreatedAt)
	if err != nil {
		return Note{}, err
	}

	updatedAt, err := time.Parse(time.RFC3339, post.UpdatedAt)
	if err != nil {
		return Note{}, err
	}
	return Note{
		ID:        post.ID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Note:      post.Note,
		UserID:    post.UserID,
	}, nil
}

func databasePostsToPosts(notes []database.Note) ([]Note, error) {
	result := make([]Note, len(notes))
	for i, note := range notes {
		var err error
		result[i], err = databaseNoteToNote(note)
		if err != nil {
			return nil, err
		}

	}
	return result, nil
}
