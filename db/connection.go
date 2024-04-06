// db/connection.go

package db

import (
	"database/sql"
	"net/url"

	_ "github.com/lib/pq"
)

func ConnectDB(serviceURI string) (*sql.DB, error) {
	conn, err := url.Parse(serviceURI)
	if err != nil {
		return nil, err
	}
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"

	db, err := sql.Open("postgres", conn.String())
	if err != nil {
		return nil, err
	}

	return db, nil
}
