package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Reservations holds the schema definition for the Reservations entity.
type Reservations struct {
	ent.Schema
}

// Fields of the Reservations.
func (Reservations) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable(),
		field.Time("created_at").Immutable(),
	}
}

// Edges of the Reservations.
func (Reservations) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("reservations").Unique(),
		edge.From("event", Events.Type).Ref("reservations").Unique(),
	}
}
