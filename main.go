package db

import (
	"context"
	"log"

	"github.com/blahcdn/db/ent/migrate"

	"github.com/blahcdn/db/ent"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *ent.Client
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
