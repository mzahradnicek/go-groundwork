package ldb

type Model struct {
	conn *Connection
}

func (m *Model) SetConnection(conn *Connection) {
	m.conn = conn
}

func (m *Model) Conn() *Connection {
	return m.conn
}

func NewModel(conn *Connection) *Model {
	return &Model{conn: conn}
}
