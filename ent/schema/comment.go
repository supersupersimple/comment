package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/rs/xid"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		// primary key
		field.String("id").
			GoType(xid.ID{}).
			DefaultFunc(xid.New),

		field.String("content"),
		field.Enum("status").Values("pending", "approved", "rejected", "deleted").Default("pending"),
		field.Time("created_at").Default(time.Now), // also treat it as publish time
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Int("depth").Default(0),

		// foreign key
		field.Int64("page_id"),
		field.Int64("user_id"),
		field.String("reply_to_id").GoType(xid.ID{}).Optional().Nillable(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("page", Page.Type).Ref("comments").Unique().Required().Field("page_id"),
		edge.From("user", User.Type).Ref("comments").Unique().Required().Field("user_id"),
		edge.To("replies", Comment.Type).
			From("reply_to").
			Unique().Field("reply_to_id"),
	}
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("page_id", "status", "reply_to_id", "id"), // for display
		index.Fields("status", "id"),                           // for admin
	}
}
