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
		field.String("host").Default("http://localhost:8080"),
		field.String("tg_bot_url").Default(""),
		field.Int("rate_limit").Default(1), // write new comment rate limit per second
	}
}

// Edges of the Conf.
func (Conf) Edges() []ent.Edge {
	return nil
}
