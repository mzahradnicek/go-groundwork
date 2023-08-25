package ldb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/blockloop/scan/v2"
	"github.com/mzahradnicek/log"
	sqlg "github.com/mzahradnicek/sql-glue"
)

var (
	ErrNoBegin    = errors.New("Can't call Begin on connection")
	ErrNoRollback = errors.New("Can't call Rollback on connection")
	ErrNoCommit   = errors.New("Can't call Commit on connection")
)

type TxBeginer interface {
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
}

type Connection struct {
	db   Querier
	sqlg *sqlg.Builder
}

func (c *Connection) GetDB() Querier {
	return c.db
}

/* SQL Glue Helpers */
func (c *Connection) GlueExec(ctx context.Context, q *sqlg.Qg) (sql.Result, error) {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return nil, log.NewError(err).AddFields(log.Fields{"query": q})
	}

	ct, err := c.db.ExecContext(ctx, sql, args...)
	if err != nil {
		err = log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	return ct, err
}

func (c *Connection) GlueExecLog(ctx context.Context, q *sqlg.Qg) (sql.Result, error) {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return nil, log.NewError(err).AddFields(log.Fields{"query": q})
	}

	ct, err := c.db.ExecContext(ctx, sql, args...)
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

	err = c.db.QueryRowContext(ctx, sql, args...).Scan(dst...)
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

	err = c.db.QueryRowContext(ctx, sql, args...).Scan(dst...)
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

	rows, err := c.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	err = scan.Rows(dst, rows)
	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	return nil
}

func (c *Connection) GlueSelectLog(ctx context.Context, q *sqlg.Qg, dst interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	rows, err := c.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	err = scan.Rows(dst, rows)
	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	log.Save(log.NewInfo("Log query").AddFields(log.Fields{"sql": sql, "args": args}))

	return nil
}

func (c *Connection) GlueGet(ctx context.Context, q *sqlg.Qg, dst interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	rows, err := c.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	err = scan.Row(dst, rows)
	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	return nil
}

func (c *Connection) GlueGetLog(ctx context.Context, q *sqlg.Qg, dst interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	row, err := c.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	err = scan.Row(dst, row)
	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
	}

	log.Save(log.NewInfo("Log query").AddFields(log.Fields{"sql": sql, "args": args}))

	return err
}

/* Transaction helpers */
func (c *Connection) Begin(ctx context.Context) (*Connection, error) {
	if v, ok := c.db.(TxBeginer); ok {
		tx, err := v.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}

		return &Connection{db: tx, sqlg: c.sqlg}, nil
	}

	return nil, ErrNoBegin
}

func (c *Connection) Rollback(ctx context.Context) error {
	if v, ok := c.db.(*sql.Tx); ok {
		return v.Rollback()
	}

	return ErrNoRollback
}

func (c *Connection) Commit(ctx context.Context) error {
	if v, ok := c.db.(*sql.Tx); ok {
		return v.Commit()
	}

	return ErrNoCommit
}
