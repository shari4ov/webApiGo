package storage

import (
	"database/sql"
	config "home/config"

	_ "github.com/lib/pq"
)

func OpenConnection() *sql.DB {
	db, err := sql.Open("postgres", config.GetPostgresConnection())
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
