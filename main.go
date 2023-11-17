package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

func main() {
	var db_name string = "./todo.db"
	// remove old sqlite db if exist
	os.Remove(db_name)

	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to database")
	}
	defer db.Close()

}
