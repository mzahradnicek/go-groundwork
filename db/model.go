package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Model struct {
	store *Connection
}

func (m *Model) SetConnection(conn *Connection) {
	m.store = conn
}
