package models

// User is a struct that contains a user's id and name
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
