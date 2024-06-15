package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Conf holds the schema definition for the Conf entity.
type Conf struct {
	ent.Schema
}

// Fields of the Conf.
func (Conf) Fields() []ent.Field {
	return []ent.Field{
		field.String("password").Sensitive(),
		field.Strings("allow_origins"),
		field.String("cookie_secret"),
		field.Int("limit_per_batch").Default(5),
		field.Int("max_loop_depth").Default(3),
	}
}

// Edges of the Conf.
func (Conf) Edges() []ent.Edge {
	return nil
}
