package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Character holds the schema definition for the Character entity.
type Character struct {
	ent.Schema
}

// Fields of the Character.
func (Character) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("character_id").
			Unique().
			Immutable(),
		field.String("name"),
		field.Uint8("faction_id"),
		field.Time("date_created"),
	}
}

// Edges of the Character.
func (Character) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("outfit", Outfit.Type).
			Ref("characters").
			Unique(),
		edge.To("ribbons", Ribbon.Type).
			Unique(),
	}

}

// Indexes of the Character.
func (Character) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("character_id"),
		index.Fields("name"),
	}
}
