package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Outfit holds the schema definition for the Outfit entity.
type Outfit struct {
	ent.Schema
}

// Fields of the Outfit.
func (Outfit) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name"),
		field.String("tag"),
		field.Uint8("faction"),
	}
}

// Edges of the Outfit.
func (Outfit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("characters", Character.Type),
	}
}
