// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/team"
	"github.com/Jecosine/alioth-kratos/app/data_service/ent/user"
)

// Team is the model entity for the Team schema.
type Team struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatedTime holds the value of the "created_time" field.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// Private holds the value of the "private" field.
	Private bool `json:"private,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TeamQuery when eager-loading is set.
	Edges        TeamEdges `json:"edges"`
	team_creator *int64
}

// TeamEdges holds the relations/edges for other nodes in the graph.
type TeamEdges struct {
	// Members holds the value of the members edge.
	Members []*User `json:"members,omitempty"`
	// Announcements holds the value of the announcements edge.
	Announcements []*Announcement `json:"announcements,omitempty"`
	// Creator holds the value of the creator edge.
	Creator *User `json:"creator,omitempty"`
	// Admins holds the value of the admins edge.
	Admins []*User `json:"admins,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) MembersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// AnnouncementsOrErr returns the Announcements value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) AnnouncementsOrErr() ([]*Announcement, error) {
	if e.loadedTypes[1] {
		return e.Announcements, nil
	}
	return nil, &NotLoadedError{edge: "announcements"}
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TeamEdges) CreatorOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.Creator == nil {
			// The edge creator was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// AdminsOrErr returns the Admins value or an error if the edge
// was not loaded in eager-loading.
func (e TeamEdges) AdminsOrErr() ([]*User, error) {
	if e.loadedTypes[3] {
		return e.Admins, nil
	}
	return nil, &NotLoadedError{edge: "admins"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Team) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case team.FieldPrivate:
			values[i] = new(sql.NullBool)
		case team.FieldID:
			values[i] = new(sql.NullInt64)
		case team.FieldName, team.FieldDescription:
			values[i] = new(sql.NullString)
		case team.FieldCreatedTime:
			values[i] = new(sql.NullTime)
		case team.ForeignKeys[0]: // team_creator
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Team", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Team fields.
func (t *Team) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case team.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int64(value.Int64)
		case team.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case team.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case team.FieldCreatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_time", values[i])
			} else if value.Valid {
				t.CreatedTime = value.Time
			}
		case team.FieldPrivate:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field private", values[i])
			} else if value.Valid {
				t.Private = value.Bool
			}
		case team.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field team_creator", value)
			} else if value.Valid {
				t.team_creator = new(int64)
				*t.team_creator = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryMembers queries the "members" edge of the Team entity.
func (t *Team) QueryMembers() *UserQuery {
	return (&TeamClient{config: t.config}).QueryMembers(t)
}

// QueryAnnouncements queries the "announcements" edge of the Team entity.
func (t *Team) QueryAnnouncements() *AnnouncementQuery {
	return (&TeamClient{config: t.config}).QueryAnnouncements(t)
}

// QueryCreator queries the "creator" edge of the Team entity.
func (t *Team) QueryCreator() *UserQuery {
	return (&TeamClient{config: t.config}).QueryCreator(t)
}

// QueryAdmins queries the "admins" edge of the Team entity.
func (t *Team) QueryAdmins() *UserQuery {
	return (&TeamClient{config: t.config}).QueryAdmins(t)
}

// Update returns a builder for updating this Team.
// Note that you need to call Team.Unwrap() before calling this method if this Team
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Team) Update() *TeamUpdateOne {
	return (&TeamClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Team entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Team) Unwrap() *Team {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Team is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Team) String() string {
	var builder strings.Builder
	builder.WriteString("Team(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", description=")
	builder.WriteString(t.Description)
	builder.WriteString(", created_time=")
	builder.WriteString(t.CreatedTime.Format(time.ANSIC))
	builder.WriteString(", private=")
	builder.WriteString(fmt.Sprintf("%v", t.Private))
	builder.WriteByte(')')
	return builder.String()
}

// Teams is a parsable slice of Team.
type Teams []*Team

func (t Teams) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
