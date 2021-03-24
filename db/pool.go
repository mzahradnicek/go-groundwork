package db

import (
	"github.com/jmoiron/sqlx"
	sqlg "github.com/mzahradnicek/sql-glue"
)

var pool map[string]*Connection

func NewConnection(ident, dbtype, connString string, sqlgBuilder *sqlg.Builder) (err error) {
	if ident == "" {
		ident = "default"
	}

	c := &Connection{}

	c.db, err = sqlx.Connect(dbtype, connString)

	if err != nil {
		return
	}

	if sqlgBuilder != nil {
		c.sqlg = sqlgBuilder
	}

	if pool == nil {
		pool = make(map[string]*Connection)
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
