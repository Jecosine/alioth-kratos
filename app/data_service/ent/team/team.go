// Code generated by entc, DO NOT EDIT.

package team

import (
	"time"
)

const (
	// Label holds the string label denoting the team type in the database.
	Label = "team"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedTime holds the string denoting the created_time field in the database.
	FieldCreatedTime = "created_time"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgeAnnouncements holds the string denoting the announcements edge name in mutations.
	EdgeAnnouncements = "announcements"
	// Table holds the table name of the team in the database.
	Table = "teams"
	// MembersTable is the table that holds the members relation/edge. The primary key declared below.
	MembersTable = "team_members"
	// MembersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	MembersInverseTable = "users"
	// AnnouncementsTable is the table that holds the announcements relation/edge.
	AnnouncementsTable = "announcements"
	// AnnouncementsInverseTable is the table name for the Announcement entity.
	// It exists in this package in order to avoid circular dependency with the "announcement" package.
	AnnouncementsInverseTable = "announcements"
	// AnnouncementsColumn is the table column denoting the announcements relation/edge.
	AnnouncementsColumn = "announcement_team"
)

// Columns holds all SQL columns for team fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedTime,
}

var (
	// MembersPrimaryKey and MembersColumn2 are the table columns denoting the
	// primary key for the members relation (M2M).
	MembersPrimaryKey = []string{"team_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedTime holds the default value on creation for the "created_time" field.
	DefaultCreatedTime time.Time
)
