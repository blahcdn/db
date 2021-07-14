package schema

import "entgo.io/ent"

// Zone holds the schema definition for the Zone entity.
type Zone struct {
	ent.Schema
}

// Fields of the Zone.
func (Zone) Fields() []ent.Field {
	return nil
}

// Edges of the Zone.
func (Zone) Edges() []ent.Edge {
	return nil
}
