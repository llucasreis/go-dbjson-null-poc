package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	connStr := "user=postgres password=postgres dbname=poc sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("[ERROR] initializing Postgres connection", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("[ERROR] pinging Postgres connection", err.Error())
		return nil, err
	}

	log.Println("[INFO] Successful connected to Postgres.")

	return db, nil
}
