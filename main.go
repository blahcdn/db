package db

import (
	"context"
	"fmt"
	"log"

	"github.com/blahcdn/db/ent/migrate"

	"github.com/blahcdn/db/ent"
	_ "github.com/lib/pq"
)

type Database struct {
	client *ent.Client
}

func Connect(ctx context.Context, url string) (*Database, error) {
	client, err := ent.Open("postgres", url)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &Database{client}, nil
}

// creates an user
func (d *Database) CreateUser(ctx context.Context, data *User) (*ent.User, error) {
	u, err := d.client.User.Create().SetEmail(data.Email).SetUsername(data.Username).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return u, nil
}

// creates a zone
func (d *Database) CreateZone(ctx context.Context, inp *Zone, ownerId int) (*ent.Zone, error) {
	z, err := d.client.Zone.Create().SetDomain(inp.Domain).SetOwnerID(ownerId).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating zone: %w", err)
	}
	return z, nil
}
