package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").Unique(),
		field.String("title").Optional(),
		field.String("url").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type),
	}
}

func (Page) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
	}
}
