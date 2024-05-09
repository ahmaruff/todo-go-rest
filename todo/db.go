package todo

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/rs/zerolog/log"
)

var (
	DB *sql.DB
)

func ConnectDatabase() error {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping error: %w", err)
	}

	DB = db
	return nil
	// db, err := sql.Open("sqlite3", "../todo.db")
	// if err != nil {
	// 	return err
	// }

	// DB = db
	// return nil
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
