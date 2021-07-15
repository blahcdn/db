package db

import (
	"context"
	"fmt"

	"github.com/blahcdn/db/ent"
	"github.com/blahcdn/db/ent/user"
	"github.com/blahcdn/db/ent/zone"
)

type Zone struct {
	Domain string
	Owner  User
}

// creates a zone
func (d *Database) CreateZone(ctx context.Context, inp *Zone, ownerId int) (*ent.Zone, error) {
	z, err := d.Client.Zone.Create().
		SetDomain(inp.Domain).
		SetOwnerID(ownerId).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating zone: %w", err)
	}
	return z, nil
}

func (d *Database) QueryZone(ctx context.Context, where *Zone) (*ent.Zone, error) {
	u, err := d.Client.Zone.
		Query().
		Where(zone.Domain(where.Domain)).
		// `Only` fails if no zone found,
		// or more than 1 zone returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying zone: %w", err)
	}
	return u, nil
}

func (d *Database) QueryZones(ctx context.Context, username string) ([]*ent.Zone, error) {
	z, err := d.Client.Zone.Query().Where(zone.HasOwnerWith(user.Username(username))).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return z, nil
}
