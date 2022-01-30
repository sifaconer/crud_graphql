package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Categories holds the schema definition for the Categories entity.
type Categories struct {
	ent.Schema
}

// Fields of the Categories.
func (Categories) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Default(""),
		field.String("description").
			NotEmpty().
			Default(""),
	}
}

// Edges of the Categories.
func (Categories) Edges() []ent.Edge {
	return nil
}
