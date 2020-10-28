package db

import (
	"github.com/jmoiron/sqlx"
	sqlg "github.com/mzahradnicek/sql-glue"
)

var pool map[string]*Connection

func NewConnection(ident, dbtype, connString string, sqlgConfig *sqlg.Config) (err error) {
	if ident == "" {
		ident = "default"
	}

	c := &Connection{}

	c.db, err = sqlx.Connect(dbtype, connString)

	if err != nil {
		return
	}

	if sqlgConfig != nil {
		c.sqlg = sqlgConfig
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
