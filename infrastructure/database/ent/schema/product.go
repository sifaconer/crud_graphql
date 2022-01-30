package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id_categories").Optional(),
		field.Float("price"),
		field.String("name").
			NotEmpty().
			Default(""),
		field.Int64("stock").Positive(),
		field.String("description").Default(""),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("categories", Categories.Type).
			Field("id_categories").
			Unique(),
	}
}
