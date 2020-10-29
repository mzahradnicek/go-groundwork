package db

type Model struct {
	conn *Connection
}

func (m *Model) SetConnection(conn *Connection) {
	m.conn = conn
}

func NewModel(conn *Connection) *Model {
	return &Model{conn: conn}
}
