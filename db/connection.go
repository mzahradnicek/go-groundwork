package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/mzahradnicek/log"
	sqlg "github.com/mzahradnicek/sql-glue"
)

type Connection struct {
	db   DbQuerier
	sqlg *sqlg.Config
	ctx  context.Context

	nestedTransaction bool
}

func (c *Connection) WithContext(ctx context.Context) *Connection {
	return &Connection{db: db, sqlg: c.sqlg, ctx: ctx, nestedTransaction: c.nestedTransaction}
}

func (c *Connection) GetDB() *sqlx.DB {
	if v, ok := c.db.(*sqlx.DB); ok {
		return v
	}
	return nil
}

func (c *Connection) GetTx(*sqlx.Tx) {
	if v, ok := c.db.(*sqlx.Tx); ok {
		return v
	}
	return nil
}

/* SQL Glue Helpers */
func (c *Connection) GlueExec(q *sqlg.Qg) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	if c.ctx != nil {
		if _, err = c.db.ExecContext(c.ctx, sql, args...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	} else {
		if _, err = c.db.Exec(sql, args...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	}

	return nil
}

func (c *Connection) GlueExecLog(q *sqlg.Qg) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	if c.ctx != nil {
		if _, err = c.db.ExecContext(c.ctx, sql, args...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	} else {
		if _, err = c.db.Exec(sql, args...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	}

	log.Save(log.NewInfo("Log query").AddFields(log.Fields{"sql": sql, "args": args}))

	return nil
}

func (c *Connection) GlueQueryRowScan(q *sqlg.Qg, dest ...interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	if c.ctx != nil {
		if err = c.db.QueryRowxContext(c.ctx, sql, args...).Scan(dest...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	} else {
		if err = c.db.QueryRowx(sql, args...).Scan(dest...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	}

	return nil
}

func (c *Connection) GlueQueryRowScanLog(q *sqlg.Qg, dest ...interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	if c.ctx != nil {
		if err = c.db.QueryRowxContext(c.ctx, sql, args...).Scan(dest...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	} else {
		if err = c.db.QueryRowx(sql, args...).Scan(dest...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	}

	log.Save(log.NewInfo("Log query").AddFields(log.Fields{"sql": sql, "args": args}))

	return nil
}

func (c *Connection) GlueSelect(q *sqlg.Qg, dest interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	if c.ctx != nil {
		if err = c.db.SelectContext(c.ctx, dest, sql, args...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	} else {
		if err = c.db.Select(dest, sql, args...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	}

	return nil
}

func (c *Connection) GlueSelectLog(q *sqlg.Qg, dest interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	if c.ctx != nil {
		if err = c.db.SelectContext(c.ctx, dest, sql, args...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	} else {
		if err = c.db.Select(dest, sql, args...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	}

	log.Save(log.NewInfo("Log query").AddFields(log.Fields{"sql": sql, "args": args}))

	return nil
}

func (c *Connection) GlueGet(q *sqlg.Qg, dest interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	if c.ctx != nil {
		if err = c.db.GetContext(c.ctx, dest, sql, args...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	} else {
		if err = c.db.Get(dest, sql, args...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	}

	return nil
}

func (c *Connection) GlueGetLog(q *sqlg.Qg, dest interface{}) error {
	sql, args, err := c.sqlg.Glue(q)

	if err != nil {
		return log.NewError(err).AddFields(log.Fields{"query": q})
	}

	if c.ctx != nil {
		if err = c.db.GetContext(c.ctx, dest, sql, args...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	} else {
		if err = c.db.Get(dest, sql, args...); err != nil {
			return log.NewError(err).AddFields(log.Fields{"sql": sql, "args": args})
		}
	}

	log.Save(log.NewInfo("Log query").AddFields(log.Fields{"sql": sql, "args": args}))

	return nil
}

/* Transaction helpers */
func (c *Connection) Begin() (*Connection, error) {
	if v, ok := c.db.(*sqlx.DB); ok {
		tx, err := v.Beginx()
		if err != nil {
			return nil, err
		}

		return &Connection{db: tx, sqlg: c.sqlg}, nil
	}

	// we already have transaction sqlx.Tx
	return &Connection{db: c.db, sqlg: c.sqlg, nestedTransaction: true}, nil
}

func (c *Connection) Rollback() error {
	if tx, ok := c.db.(*sqlx.Tx); ok && !c.nestedTransaction {
		return tx.Rollback()
	}

	return nil
}

func Commit() error {
	if tx, ok := c.db.(*sqlx.Tx); ok && !c.nestedTransaction {
		return tx.Commit()
	}

	return nil
}
