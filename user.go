package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/blahcdn/db/ent"
	"github.com/blahcdn/db/ent/user"
)

type User struct {
	LowerUsername string
	PasswordHash  []byte
	Username      string
	Email         string
	Zones         []Zone
}

// creates an user
func (d *Database) CreateUser(ctx context.Context, data *User) (*ent.User, error) {
	u, err := d.Client.User.Create().SetEmail(data.Email).SetUsername(data.Username).SetLowerUsername(strings.ToLower(data.Username)).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return u, nil
}

func (d *Database) QueryUser(ctx context.Context, where *User) (*ent.User, error) {
	u, err := d.Client.User.
		Query().
		Where(user.LowerUsername(strings.ToLower(where.Username))).
		Where(user.Email(where.Email)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return u, nil
}
