package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

var (
	ErrCredentialNotMatched    = errors.New("credential not matched")
	ErrIncompatibleHashVersion = errors.New("incompatible hash version")
)

const (
	saltLength     = 16
	hashIterations = 3
	memory         = 64 * 1024
	parallelism    = 4
	keyLength      = 32
)

type Credentials struct {
	Hash          string   `json:"-"`
	SshPublicKeys []string `json:"-"`
}

func (c Credentials) ValidatePassword(password string) error {
	var (
		version    int
		mem        uint32
		iterations uint32
		parallel   uint8
	)

	vals := strings.SplitN(c.Hash, "$", 6)

	_, err := fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return err
	}

	if version != argon2.Version {
		return ErrIncompatibleHashVersion
	}

	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &mem, &iterations, &parallel)
	if err != nil {
		return err
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return err
	}

	hash, err := base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return err
	}

	reHash := argon2.IDKey([]byte(password), salt, iterations, mem, parallel, keyLength)

	if subtle.ConstantTimeCompare(hash, reHash) != 1 {
		return ErrCredentialNotMatched
	}

	return nil
}

func NewPasswordCredentials(password string) (Credentials, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return Credentials{}, nil
	}

	hash := argon2.IDKey([]byte(password), salt, hashIterations, memory, parallelism, keyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		memory,
		hashIterations,
		parallelism,
		b64Salt,
		b64Hash,
	)

	return Credentials{
		Hash: encodedHash,
	}, nil
}
