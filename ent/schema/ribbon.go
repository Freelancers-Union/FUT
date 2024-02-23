package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Ribbon holds the schema definition for the Ribbon entity.
type Ribbon struct {
	ent.Schema
}

// Fields of the Ribbon.
func (Ribbon) Fields() []ent.Field {
	return []ent.Field{
		field.Time("timestamp"),
		field.Int("ribbon_id"),
		field.Int("ribbon_count"),
	}
}

// Edges of the Ribbon.
func (Ribbon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("character", Character.Type).
			Ref("ribbons").
			Unique(),
	}
}

// // Annotations of the Ribbon.
// func (Ribbon) Annotations() []schema.Annotation {
// 	return []schema.Annotation{
// 		ent.PrimaryKey("timestamp", "character_id", "ribbon_id"),
// 	}
// }
//
