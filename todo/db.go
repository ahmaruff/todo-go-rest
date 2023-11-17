package todo

import (
	"database/sql"
	"errors"
	"os"

	"github.com/rs/zerolog/log"
)

var (
	db *sql.DB
)

func SetDatabase(db_name string) error {
	var err error

	if db_name == "" {
		return errors.New("There is no database")
	}

	// remove old sqlite db if exist
	os.Remove(db_name)

	db, err = sql.Open("sqlite3", db_name)
	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to database")
	}
	defer db.Close()

	return nil
}
