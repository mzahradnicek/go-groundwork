package db

import (
	"errors"

	"github.com/jmoiron/sqlx"
	sqlg "github.com/mzahradnicek/sql-glue"
)

var pool = make(map[string]*Connection)

var ErrDbIsNil = errors.New("Database identifier is nil")

func NewConnection(ident, dbtype, connString string, sqlgBuilder *sqlg.Builder) (err error) {
	db, err := sqlx.Connect(dbtype, connString)

	if err != nil {
		return err
	}

	return NewConnectionFromDb(ident, db, sqlgBuilder)
}

func NewConnectionFromDb(ident string, db *sqlx.DB, sqlgBuilder *sqlg.Builder) (err error) {
	if ident == "" {
		ident = "default"
	}

	if db == nil {
		return ErrDbIsNil
	}

	c := &Connection{db: db}

	if sqlgBuilder != nil {
		c.sqlg = sqlgBuilder
	}

	pool[ident] = c

	return
}

func GetConnection(name ...string) *Connection {
	n := "default"
	if len(name) > 0 {
		n = name[0]
	}

	return pool[n]
}
