package auth

import (
	"context"
	"fmt"
	"github.com/NV4RE/clarchgo/entity/auth"
	authRepo "github.com/NV4RE/clarchgo/repository/auth"
	"time"
)

type UseCase struct {
	authRepo authRepo.Repository
}

func (uc UseCase) CreateUserWithPassword(ctx context.Context, usr auth.User, username, password string) (auth.User, error) {
	err := usr.IsAllowed(auth.PermissionAuthUserCompanyWrite)
	if err != nil {
		return auth.User{}, err
	}

	cred, err := auth.NewPasswordCredentials(password)
	if err != nil {
		return auth.User{}, err
	}

	u := auth.NewUser(username, cred)
	err = uc.authRepo.CreateUser(ctx, u)

	return u, err
}

func (uc UseCase) GetUserByUsername(ctx context.Context, usr auth.User, username string) (auth.User, error) {
	err := usr.IsAllowed(auth.PermissionAuthUserCompanyRead)
	if err != nil {
		return auth.User{}, err
	}

	return uc.authRepo.GetUserByUsername(ctx, username)
}

func (uc UseCase) ListAllUsers(ctx context.Context, usr auth.User) ([]auth.User, error) {
	err := usr.IsAllowed(auth.PermissionAuthUserCompanyRead)
	if err != nil {
		return []auth.User{}, err
	}

	return uc.authRepo.ListUsers(ctx)
}

func (uc UseCase) ListActivatedUsers(ctx context.Context, usr auth.User) ([]auth.User, error) {
	err := usr.IsAllowed(auth.PermissionAuthUserCompanyRead)
	if err != nil {
		return []auth.User{}, err
	}

	return uc.authRepo.ListUsersByIsActivated(ctx, true)
}

func (uc UseCase) ListNotActivatedUsers(ctx context.Context, usr auth.User) ([]auth.User, error) {
	err := usr.IsAllowed(auth.PermissionAuthUserCompanyRead)
	if err != nil {
		return []auth.User{}, err
	}

	return uc.authRepo.ListUsersByIsActivated(ctx, false)
}

func (uc UseCase) ActivatedUserByUsername(ctx context.Context, usr auth.User, username string) error {
	err := usr.IsAllowed(auth.PermissionAuthUserCompanyWrite)
	if err != nil {
		return err
	}

	return uc.authRepo.ActivateUser(ctx, username)
}

func (uc UseCase) DeleteUserByUsername(ctx context.Context, usr auth.User, username string) error {
	err := usr.IsAllowed(auth.PermissionAuthUserCompanyWrite)
	if err != nil {
		return err
	}

	return uc.authRepo.DeleteUserByUsername(ctx, username)
}

func (uc UseCase) LoginWithPassword(ctx context.Context, username, password string) error {
	u, err := uc.authRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	err = u.Credentials.ValidatePassword(password)
	if err != nil {
		return err
	}

	err = uc.authRepo.CreateToken(
		ctx,
		fmt.Sprintf("%s-%s", username, time.Now().Format(time.RFC3339)),
		u,
		time.Now().Add(365*24*time.Hour),
	)

	return err
}

func (uc UseCase) GetUserByToken(ctx context.Context, token string) (auth.User, error) {
	return uc.authRepo.GetUserByToken(ctx, token)
}

func NewUseCase(authRepo authRepo.Repository) *UseCase {
	return &UseCase{
		authRepo: authRepo,
	}
}
