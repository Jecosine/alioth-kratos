package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("text").NotEmpty(),
		field.Bool("done").Default(false),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
