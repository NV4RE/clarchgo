package auth

import (
	"context"
	"fmt"
	"github.com/NV4RE/clarchgo/entity/auth"
	"testing"
	"time"
)

func TestMongoDB_Case1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	m, err := NewMongo("mongodb://localhost:27017/", fmt.Sprintf("test-auth-%s", time.Now().Format(time.RFC3339)))
	exitIfError(err, t)

	userNames := []string{"user1", "user2", "user3", "user4", "user5"}

	for _, un := range userNames {
		cred, err := auth.NewPasswordCredentials("password")
		if err != nil {
			t.Error(err)
			return
		}
		u := auth.NewUser(un, cred)

		err = m.CreateUser(ctx, u)
		if err != nil {
			t.Error(err)
			return
		}
	}

	users, err := m.ListUsers(ctx)
	exitIfError(err, t)

	if len(users) != len(userNames) {
		t.Error("user not matched")
	}

	for _, un := range userNames {
		u, err := m.GetUserByUsername(ctx, un)
		if err != nil {
			t.Error(err)
			return
		}

		if u.Username != un {
			t.Error("username not matched")
			return
		}
	}

	cred, err := auth.NewPasswordCredentials("password")
	exitIfError(err, t)
	u := auth.NewUser(userNames[1], cred)

	err = m.CreateUser(ctx, u)
	if err != ErrUserAlreadyExist {
		t.Error("should not able to create user", err)
		return
	}

	users, err = m.ListUsersByIsActivated(ctx, true)
	exitIfError(err, t)

	if len(users) != 0 {
		t.Error("should not found any activated users", err)
		return
	}

	err = m.ActivateUser(ctx, userNames[0])
	exitIfError(err, t)

	users, err = m.ListUsersByIsActivated(ctx, true)
	exitIfError(err, t)

	if len(users) != 1 {
		t.Error("should only found 1 activated user", err)
		return
	}

	if users[0].Username != userNames[0] {
		t.Error("activated wrong user", err)
		return
	}

	users, err = m.ListUsers(ctx)
	exitIfError(err, t)

	if len(users) != len(userNames) {
		t.Error("user not matched")
		return
	}

	// Delete user and re-create that user

	err = m.DeleteUserByUsername(ctx, userNames[1])
	exitIfError(err, t)

	users, err = m.ListUsers(ctx)
	exitIfError(err, t)

	if len(users) != len(userNames)-1 {
		t.Error("user not matched")
		return
	}

	u, err = m.GetUserByUsername(ctx, userNames[1])
	if err != ErrUserNotFound {
		t.Error("user should be deleted")
		return
	}

	cred, err = auth.NewPasswordCredentials("password")
	exitIfError(err, t)
	u = auth.NewUser(userNames[1], cred)

	err = m.CreateUser(ctx, u)
	if err != nil {
		t.Error(err)
		return
	}

	err = m.CreateUser(ctx, u)
	if err != ErrUserAlreadyExist {
		t.Error("should not able to create user", err)
		return
	}
}

func exitIfError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
