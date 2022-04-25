package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("nickname").NotEmpty(),
		field.String("email"),
		field.String("password").NotEmpty(),
		field.String("avatar").Default("test").NotEmpty(),
		field.Time("created_time").Default(time.Now()).Immutable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teams", Team.Type).
			Annotations(entgql.Bind()).Ref("members").Annotations(entgql.Bind()),
		edge.From("announcements", Announcement.Type).
			Annotations(entgql.Bind()).
			Ref("author").
			Annotations(entgql.Bind()),
		edge.From("records", JudgeRecord.Type).
			Annotations(entgql.Bind()).
			Ref("user").
			Annotations(entgql.Bind()),
		edge.To("created_problems", Problem.Type).
			Annotations(entgql.Bind()).
			From("author"),
		edge.To("solved_problems", Problem.Type).
			Annotations(entgql.Bind()).
			From("solved_by").
			Annotations(entgql.Bind()),
	}
}
