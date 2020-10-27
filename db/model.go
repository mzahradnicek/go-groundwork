package db

type Model struct {
	store *Connection
}

func (m *Model) SetConnection(conn *Connection) {
	m.store = conn
}
