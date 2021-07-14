package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blahcdn/db/ent"
	_ "github.com/lib/pq"
)

type Database struct {
	client *ent.Client
}

type User struct {
	Username string
	Email    string
}

func Connect(url string) (*Database, error) {
	client, err := ent.Open("postgres", url)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &Database{client}, nil
}

func (d *Database) CreateUser(ctx context.Context) (*ent.User, error) {
	u, err := d.client.User.
		Create().
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}
