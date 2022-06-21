package db

import (
	"bytes"
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func url(conf Storage) string {
	var b bytes.Buffer

	b.WriteString("postgres://")
	b.WriteString(conf.User)
	b.WriteString(":")
	b.WriteString(conf.Password)
	b.WriteString("@")
	b.WriteString(conf.Host)
	b.WriteString(":")
	b.WriteString(conf.Port)
	b.WriteString("/")
	b.WriteString(conf.Name)
	b.WriteString("?sslmode=disable")

	return b.String()
}

// Read instance connection to postgresql
func (conf Storage) Conn() (*sql.DB, error) {
	conn, err := sql.Open("postgres", url(conf))
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, err
}
