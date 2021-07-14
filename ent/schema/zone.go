package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Zone holds the schema definition for the Zone entity.
type Zone struct {
	ent.Schema
}

// Fields of the Zone.
func (Zone) Fields() []ent.Field {
	return []ent.Field{field.String("domain").Unique()}
}

// Edges of the Zone.
func (Zone) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("zones").Unique(),
	}
}
