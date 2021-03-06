package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("name").Unique(),
		field.String("description"),
		field.Time("created_time").Default(time.Now()).Immutable(),
		field.Bool("private").Default(false),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", User.Type).Annotations(entgql.Bind()),
		edge.From("announcements", Announcement.Type).Annotations(entgql.Bind()).Ref("team").Annotations(entgql.Bind()),
		edge.To("creator", User.Type).Annotations(entgql.Bind()).Unique(),
		edge.To("admins", User.Type).Annotations(entgql.Bind()),
	}
}
