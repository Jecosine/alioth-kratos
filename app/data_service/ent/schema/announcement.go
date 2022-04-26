package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Announcement holds the schema definition for the Announcement entity.
type Announcement struct {
	ent.Schema
}

// Fields of the Announcement.
func (Announcement) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("title").NotEmpty(),
		field.String("content"),
		field.Time("createdTime").Default(time.Now()).Immutable(),
	}
}

// Edges of the Announcement.
func (Announcement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).Annotations(entgql.Bind()),
	}
}
