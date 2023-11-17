package todo

import (
	"database/sql"

	"github.com/rs/zerolog/log"
)

var (
	db *sql.DB
)

func SetDatabase() error {
	var err error

	var db_name string = "../todo.db"

	db, err = sql.Open("sqlite3", db_name)
	if err != nil {
		log.Error().Err(err).Msg("unable to connect to database")
	}

	log.Info().Any("db", db)
	query := `CREATE TABLE IF NOT EXISTS todolist (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		done_at DATETIME,
		is_done BOOLEAN GENERATED ALWAYS AS (done_at IS NOT NULL) STORED
	);`

	_, err = db.Exec(query)

	log.Info().Any("db", db)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create database")
		return err
	}

	return nil
}

func getDatabase() *sql.DB {
	return db
}