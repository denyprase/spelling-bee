package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Init(sourceName string) *sql.DB {
	conn, err := sql.Open("pgx", sourceName)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	return conn
}
