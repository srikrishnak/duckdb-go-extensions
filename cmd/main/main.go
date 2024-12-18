package main

import (
	"database/sql"
	"github.com/marcboeker/go-duckdb"
	duckdb_go_extensions "github.com/sundeck-io/duckdb-go-extensions"
)

func main() {
	conn, _ := duckdb.NewConnector("", nil)

	db := sql.OpenDB(conn)
	db.Ping()

	duckdb_go_extensions.LoadExtensions(conn)
}
