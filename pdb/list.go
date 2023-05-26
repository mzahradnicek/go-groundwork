package pdb

import (
	sqlg "github.com/mzahradnicek/sql-glue"
)

type List struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`

	conn *Connection

	opt    *QueryOptions
	params QueryParams
	where  *sqlg.Qg // DEPRECATED
	filter sqlg.Qg
}

func (l *List) SetConnection(conn *Connection) {
	l.conn = conn
}

func (l List) Conn() *Connection {
	return l.conn
}

func (l *List) SetFilter(qg sqlg.Qg) {
	l.filter = qg
}

func (l List) GetFilter() sqlg.Qg {
	return l.filter
}

func (l *List) SetParams(p QueryParams) {
	l.params = p
}

func (l List) GetParams() QueryParams {
	return l.params
}

// DEPRECATED
func (l *List) SetWhere(qg *sqlg.Qg) {
	l.where = qg
}

// DEPRECATED
func (l List) GetWhere() *sqlg.Qg {
	return l.where
}

func (l *List) QueryOptions(opt QueryOptions) {
	l.opt = &opt
}

func (l *List) ApplyQueryOptions(q *sqlg.Qg) {
	if l.opt == nil {
		return
	}

	l.opt.ApplyToQuery(q)
	l.Limit = l.opt.Limit
	l.Offset = l.opt.Offset
}

func NewList(conn *Connection) *List {
	res := &List{conn: conn}
	res.params = make(QueryParams)
	return res
}
