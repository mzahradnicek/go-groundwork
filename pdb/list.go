package pdb

import (
	sqlg "github.com/mzahradnicek/sql-glue"
)

type List struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`

	conn *Connection

	opt   *QueryOptions
	where *sqlg.Qg
}

func (l *List) SetConnection(conn *Connection) {
	l.conn = conn
}

func (l List) Conn() *Connection {
	return l.conn
}

func (l *List) SetWhere(qg sqlg.Qg) {
	l.where = &qg
}

func (l List) GetWhere() sqlg.Qg {
	return *l.where
}

func (l *List) QueryOptions(opt QueryOptions) {
	l.opt = &opt
}

func (l List) ApplyQueryOptions(q *sqlg.Qg) {
	if l.opt != nil {
		l.opt.ApplyToQuery(q)

		l.Limit = l.opt.Limit
		l.Offset = l.opt.Offset
	}
}

func NewList(conn *Connection) *List {
	return &List{conn: conn}
}
