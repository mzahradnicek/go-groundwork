package models

import (
	. "alcrm/common"

	"context"

	"github.com/jmoiron/sqlx"
)

type Model struct {
	store *Connection
}

func (m *Model) SetConnection(conn *Connection) {
	m.store = conn
	if m.ctx != nil {
		m.store = conn.WithContext(m.ctx)
	}
}
