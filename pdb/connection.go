package pdb

import (
	"context"
	"errors"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mzahradnicek/log"
	sqlg "github.com/mzahradnicek/sql-glue"
)

var (
	ErrNoBegin    = errors.New("Can't call Begin on connection")
	ErrNoRollback = errors.New("Can't call Rollback on connection")
	ErrNoCommit   = errors.New("Can't call Commit on connection")
)

type Connection struct {
	db   Querier
	sqlg *sqlg.Builder
}

func (c *Connection) GetDB() Querier {
	return c.db
}

/* SQL Glue Helpers */
func (c *Connection) GlueExec(ctx context.Context, q *sqlg.Qg) (pgconn.CommandTag, error) {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return pgconn.CommandTag{}, log.NewError(err).AddFields(log.Fields{"query": q})
	}

	ct, err := c.db.Exec(ctx, sql, args...)
	if err != nil {
		err = log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	return ct, err
}

func (c *Connection) GlueExecLog(ctx context.Context, q *sqlg.Qg) (pgconn.CommandTag, error) {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return pgconn.CommandTag{}, log.NewError(err).AddFields(log.Fields{"query": q})
	}

	ct, err := c.db.Exec(ctx, sql, args...)
	if err != nil {
		err = log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	log.Save(log.NewInfo("Log query").AddFields(log.Fields{"sql": sql, "args": args}))

	return ct, err
}

func (c *Connection) GlueQueryRowScan(ctx context.Context, q *sqlg.Qg, dst ...interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	err = c.db.QueryRow(ctx, sql, args...).Scan(dst...)
	if err != nil {
		err = log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	return err
}

func (c *Connection) GlueQueryRowScanLog(ctx context.Context, q *sqlg.Qg, dst ...interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	err = c.db.QueryRow(ctx, sql, args...).Scan(dst...)
	if err != nil {
		err = log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	log.Save(log.NewInfo("Log query").AddFields(log.Fields{"sql": sql, "args": args}))

	return err
}

func (c *Connection) GlueSelect(ctx context.Context, q *sqlg.Qg, dst interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	err = pgxscan.Select(ctx, c.db, dst, sql, args...)
	if err != nil {
		err = log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	return err
}

func (c *Connection) GlueSelectLog(ctx context.Context, q *sqlg.Qg, dst interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	err = pgxscan.Select(ctx, c.db, dst, sql, args...)
	if err != nil {
		err = log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	log.Save(log.NewInfo("Log query").AddFields(log.Fields{"sql": sql, "args": args}))

	return err
}

func (c *Connection) GlueGet(ctx context.Context, q *sqlg.Qg, dst interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	err = pgxscan.Get(ctx, c.db, dst, sql, args...)
	if err != nil {
		err = log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	return err
}

func (c *Connection) GlueGetLog(ctx context.Context, q *sqlg.Qg, dst interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	err = pgxscan.Get(ctx, c.db, dst, sql, args...)
	if err != nil {
		err = log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	log.Save(log.NewInfo("Log query").AddFields(log.Fields{"sql": sql, "args": args}))

	return err
}

/* Transaction helpers */
func (c *Connection) Begin(ctx context.Context) (*Connection, error) {
	if v, ok := c.db.(pgx.Tx); ok {
		tx, err := v.Begin(ctx)
		if err != nil {
			return nil, err
		}

		return &Connection{db: tx, sqlg: c.sqlg}, nil
	}

	return nil, ErrNoBegin
}

func (c *Connection) Rollback(ctx context.Context) error {
	if v, ok := c.db.(pgx.Tx); ok {
		return v.Rollback(ctx)
	}

	return ErrNoRollback
}

func (c *Connection) Commit(ctx context.Context) error {
	if v, ok := c.db.(pgx.Tx); ok {
		return v.Commit(ctx)
	}

	return ErrNoCommit
}
