package db

import (
	"context"

	sqlg "github.com/mzahradnicek/sql-glue"
)

type List struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`

	store DbQuerier

	opt   *QueryOptions
	where sqlg.Qg
}

func (l *List) SetStore(dbh DbQuerier) {
	l.store = dbh
}

func (l *List) QueryGlue(q sqlg.Qg) {
	l.where = q
}

func (l *List) QueryOptions(opt QueryOptions) {
	l.opt = &opt
}
