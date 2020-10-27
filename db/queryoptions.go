package common

import (
	"errors"
	"net/url"
	"strconv"
	"strings"

	"github.com/mzahradnicek/go-groundwork/utils"
	sqlg "github.com/mzahradnicek/sql-glue"
)

var (
	ErrNoLimitDefined = errors.New("No limit defined")
)

type QueryOptions struct {
	Limit       int
	Offset      int
	Page        int
	Sort        []string
	SortAllow   []string
	SortDefault string
}

func (qOpts *QueryOptions) GetFromMap(m map[string]string) {
	if v, ok := m["limit"]; ok {
		qOpts.Limit, _ = strconv.Atoi(v)
	}

	if v, ok := m["offset"]; ok {
		qOpts.Offset, _ = strconv.Atoi(v)
	}

	if v, ok := m["page"]; ok {
		qOpts.Page, _ = strconv.Atoi(v)
	}
}

func (qOpts *QueryOptions) GetFromURLQuery(m url.Values) {
	if v, ok := m["limit"]; ok {
		qOpts.Limit, _ = strconv.Atoi(v[0])
	}

	if v, ok := m["offset"]; ok {
		qOpts.Offset, _ = strconv.Atoi(v[0])
	}

	if v, ok := m["page"]; ok {
		qOpts.Page, _ = strconv.Atoi(v[0])
	}

	if v, ok := m["sort"]; ok {
		for _, o := range v {
			s := strings.Split(o, ":")
			if !utils.SliceStringContains(qOpts.SortAllow, s[0]) {
				continue
			}

			if len(s) == 2 && s[1] == "desc" {
				s[0] += " DESC"
			}

			qOpts.Sort = append(qOpts.Sort, s[0])
		}
	}
}

func (opt *QueryOptions) ApplyToQuery(sb *sqlg.Qg) error {
	if opt.Page > 0 {
		if opt.Limit == 0 {
			return ErrNoLimitDefined
		}

		opt.Offset = opt.Limit * (opt.Page - 1)
	}

	if len(opt.Sort) > 0 {
		sb.Append("ORDER BY", strings.Join(opt.Sort, ", "))
	} else if opt.SortDefault != "" {
		sb.Append("ORDER BY", opt.SortDefault)
	}

	if opt.Limit > 0 {
		sb.Append("LIMIT %v", opt.Limit)
	}

	if opt.Offset > 0 {
		sb.Append("OFFSET %v", opt.Offset)
	}

	return nil
}
