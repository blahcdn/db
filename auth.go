package db

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"fmt"

	"strings"

	"golang.org/x/crypto/argon2"
)

func argon2IDKey(input []byte, salt []byte) []byte {
	return argon2.IDKey(input, salt, 1, 64*1024, 4, 32)
}

func PasswordHash(plaintext string) ([]byte, error) {
	// Generate a random 16-byte salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return make([]byte, 0), err
	}

	hash := argon2IDKey([]byte(plaintext), salt)

	return append(salt, hash...), nil
}

func ValidateHash(payload []byte, plaintext string) bool {
	salt := payload[:16]
	hash := payload[16:]

	providedHash := argon2IDKey([]byte(plaintext), salt)
	return subtle.ConstantTimeCompare(hash, providedHash) == 1
}

func (d *Database) AuthorizeUser(ctx context.Context, username string, plaintext string) (bool, error) {
	u, err := d.QueryUser(ctx, &User{LowerUsername: strings.ToLower(username)})
	if err != nil {
		return false, fmt.Errorf("failed validating password: %w", err)
	}
	return ValidateHash(u.PasswordHash, plaintext), nil

}
