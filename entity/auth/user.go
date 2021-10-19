package auth

import "time"

type User struct {
	Username    string       `json:"username"`
	Credentials *Credentials `json:"-"`
	Roles       []Role       `json:"roles"`
	IsActivated bool         `json:"is_activated"`
	CreatedAt   time.Time    `json:"created_at"`
}

func NewUser(username string, cred Credentials) User {
	return User{
		Username:    username,
		Credentials: &cred,
		Roles:       make([]Role, 0),
		IsActivated: false,
		CreatedAt:   time.Now(),
	}
}
