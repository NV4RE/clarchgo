package auth

import (
	"context"
	"errors"
	"github.com/NV4RE/clarchgo/entity/auth"
	"time"
)

var (
	ErrUserAlreadyExist     = errors.New("user already exist")
	ErrUserNotFound         = errors.New("user not found")
	ErrTokenCannotBeCreated = errors.New("token cannot be created")
)

type Repository interface {
	CreateUser(ctx context.Context, user auth.User) error
	GetUserByUsername(ctx context.Context, username string) (auth.User, error)
	ListUsers(ctx context.Context) ([]auth.User, error)
	ListUsersByIsActivated(ctx context.Context, isActivated bool) ([]auth.User, error)
	ActivateUser(ctx context.Context, username string) error
	DeleteUserByUsername(ctx context.Context, username string) error

	CreateToken(ctx context.Context, token string, user auth.User, expiresBy time.Time) error
	GetUserByToken(ctx context.Context, token string) (auth.User, error)
}
