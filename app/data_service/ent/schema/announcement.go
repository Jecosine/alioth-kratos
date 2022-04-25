package schema

import "entgo.io/ent"

// Announcement holds the schema definition for the Announcement entity.
type Announcement struct {
	ent.Schema
}

// Fields of the Announcement.
func (Announcement) Fields() []ent.Field {
	return nil
}

// Edges of the Announcement.
func (Announcement) Edges() []ent.Edge {
	return nil
}
