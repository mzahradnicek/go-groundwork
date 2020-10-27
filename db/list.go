package db

import (
	sqlg "github.com/mzahradnicek/sql-glue"
)

type List struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`

	store *Connection

	opt   *QueryOptions
	where sqlg.Qg
}

func (l *List) SetConnection(conn *Connection) {
	l.store = conn
}

func (l *List) QueryGlue(q sqlg.Qg) {
	l.where = q
}

func (l *List) QueryOptions(opt QueryOptions) {
	l.opt = &opt
}
