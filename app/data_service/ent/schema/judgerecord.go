package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JudgeRecord holds the schema definition for the JudgeRecord entity.
type JudgeRecord struct {
	ent.Schema
}

// Fields of the JudgeRecord.
func (JudgeRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Time("judge_time"),
		field.Time("finished_time"),
		field.Int64("time_cost"),
		field.Int64("memory_cost"),
		field.Int64("status").Default(0),
	}
}

// Edges of the JudgeRecord.
func (JudgeRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Annotations(entgql.Bind()),
	}
}
