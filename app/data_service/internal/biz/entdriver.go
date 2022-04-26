package biz

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
)

func NewEntDriver(db *sql.DB) *entsql.Driver {
	return entsql.OpenDB(dialect.Postgres, db)
}
