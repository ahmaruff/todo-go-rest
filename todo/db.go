package todo

import (
	"database/sql"

	"github.com/rs/zerolog/log"
)

var (
	DB *sql.DB
)

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "../todo.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func SetupDatabase() {
	err := ConnectDatabase()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to database")
	}

	query := `CREATE TABLE IF NOT EXISTS todolist (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		done_at DATETIME,
		is_done BOOLEAN GENERATED ALWAYS AS (done_at IS NOT NULL) STORED
	);`

	_, err = DB.Exec(query)

	if err != nil {
		log.Fatal().Err(err).Msg("unable to create table database")
	}

	// defer DB.Close()
}
