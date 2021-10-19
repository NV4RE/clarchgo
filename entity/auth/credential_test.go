package auth

import (
	"fmt"
	"testing"
)

func TestNewPasswordCredentials(t *testing.T) {
	password := "super-secret"
	{
		cred, err := NewPasswordCredentials(password)
		fmt.Println(cred.Hash)
		if err != nil {
			t.Error(err)
		}

		if cred.Hash == "" {
			t.Error("hash should not be empty")
		}

		err = cred.ValidatePassword(password)
		if err != nil {
			t.Error(err)
		}

		err = cred.ValidatePassword("")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword(" super-secret")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword(" super-secret ")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword("super-secret😋")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword("😂😃🧘🏻‍♂🌍🍞🚗📞🎉❤️🍆")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}
	}

	password = "😂😃🧘🏻‍♂🌍🍞🚗📞🎉❤️🍆"
	{
		cred, err := NewPasswordCredentials(password)
		fmt.Println(cred.Hash)
		if err != nil {
			t.Error(err)
		}

		if cred.Hash == "" {
			t.Error("hash should not be empty")
		}

		err = cred.ValidatePassword(password)
		if err != nil {
			t.Error(err)
		}

		err = cred.ValidatePassword("")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword(" super-secret")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword(" super-secret ")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword("super-secret😋")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}
	}

}

// Backward compatible test
func TestCredentials_ValidatePassword(t *testing.T) {
	{
		cred := Credentials{
			Hash: "$argon2id$v=19$m=65536,t=3,p=4$+Kc2KbSzeQckbuJiSEpkpw$bsQU4n6AHYhsk9aAh4TUnSpggPc6KjaQ6GMeWse1RVU",
		}

		err := cred.ValidatePassword("super-secret")
		if err != nil {
			t.Error(err)
		}

		err = cred.ValidatePassword("")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword(" super-secret")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword(" super-secret ")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword("super-secret😋")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword("😂😃🧘🏻‍♂🌍🍞🚗📞🎉❤️🍆")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}
	}

	{
		cred := Credentials{
			Hash: "$argon2id$v=19$m=65536,t=3,p=4$wot/BbBpw7RTdfl63ylWDw$/CsWG+2Cj3SddBzPiN5cCjrnLIj/Vg1rT4nWNOx8d0Y",
		}

		err := cred.ValidatePassword("😂😃🧘🏻‍♂🌍🍞🚗📞🎉❤️🍆")
		if err != nil {
			t.Error(err)
		}

		err = cred.ValidatePassword("")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword(" super-secret")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword(" super-secret ")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword("super-secret😋")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}

		err = cred.ValidatePassword(" 😂😃🧘🏻‍♂🌍🍞🚗📞🎉❤️🍆")
		if err != ErrCredentialNotMatched {
			t.Errorf("should return with error ErrCredentialNotMatched")
		}
	}
}
