package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive().Default(1),
		field.String("name").
			Default("unknown"),
		field.String("email").Unique(),
		field.String("password"),
		field.String("username").Unique(),
		field.Bool("is_admin").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Events.Type).
			Ref("users"),
		edge.To("reservations", Reservations.Type),
	}
}
