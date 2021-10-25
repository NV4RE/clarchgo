package auth

import (
	"errors"
	"time"
)

var (
	ErrUserDoesNotHaveAnyPermission = errors.New("user does not have any permission")
)

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

func (u User) IsAllowed(perm Permission) error {
	if len(u.Roles) == 0 {
		return ErrUserDoesNotHaveAnyPermission
	}

	for _, r := range u.Roles {
		err := r.IsAllowed(perm)
		if err != nil {
			return err
		}
	}

	return nil
}
