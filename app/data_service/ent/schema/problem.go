package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Problem holds the schema definition for the Problem entity.
type Problem struct {
	ent.Schema
}

// Fields of the Problem.
func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("title").NotEmpty(),
		field.String("content"),
		field.Time("created_time").Default(time.Now()).Immutable(),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Annotations(entgql.Bind()).Ref("created_problems").Annotations(entgql.Bind()).Unique(),
		edge.From("solved_by", User.Type).Annotations(entgql.Bind()).Ref("solved_problems").Annotations(entgql.Bind()),
		edge.To("tags", Tag.Type).Annotations(entgql.Bind()),
	}
}
