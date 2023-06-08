package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

func (m *DB) Ping() error {
	return m.conn.Ping()
}

func (m *DB) Close() error {
	return m.conn.Close()
}

func (m *DB) GetConn() *sql.DB {
	return m.conn
}

func New(connStr string) (*DB, error) {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &DB{conn: conn}, nil
}
