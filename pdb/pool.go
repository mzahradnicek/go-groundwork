package pdb

import (
	"context"
	"errors"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlg "github.com/mzahradnicek/sql-glue"
)

var pool = make(map[string]*Connection)

var ErrDbIsNil = errors.New("Database identifier is nil")

func NewConnection(ident, connString string, sqlgBuilder *sqlg.Builder) error {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return err
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxdecimal.Register(conn.TypeMap())
		return nil
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		return err
	}

	return NewConnectionFromDb(ident, db, sqlgBuilder)
}

func NewConnectionFromDb(ident string, db Querier, sqlgBuilder *sqlg.Builder) error {
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

	return nil
}

func GetConnection(name ...string) *Connection {
	n := "default"
	if len(name) > 0 {
		n = name[0]
	}

	return pool[n]
}
