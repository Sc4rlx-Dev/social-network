package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	Nickname     string    `json:"nickname"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	PasswordHash string    `json:"-"` // Omit from JSON responses
	DateOfBirth  string    `json:"date_of_birth"`
	Gender       string    `json:"gender"`
	CreatedAt    time.Time `json:"created_at"`
}