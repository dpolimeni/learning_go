package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Events holds the schema definition for the Events entity.
type Events struct {
	ent.Schema
}

// Fields of the Events.
func (Events) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").Positive(),
		field.String("name").Default("unknown"),
		field.Int16("capacity").Positive().Max(10000),
		field.String("description").Default("unknown"),
	}
}

// Edges of the Events.
func (Events) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
