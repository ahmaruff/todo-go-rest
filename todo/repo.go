package todo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

var ErrTodoNotFound error = errors.New("todo: not found")

func findItemById(ctx context.Context, db *sql.DB, id ulid.ULID) (TodoItem, error) {
	q := `SELECT id, title, created_at, done_at FROM todolist WHERE id=$1`

	// row, err := db.ExecContext(ctx, q, id)
	row := db.QueryRowContext(ctx, q, id)

	var item TodoItem
	if err := row.Scan(&item.Id, &item.Title, &item.CreatedAt, &item.DoneAt); err != nil {
		if err == sql.ErrNoRows {
			log.Debug().Err(err).Msg("Record not found")
			return TodoItem{}, ErrTodoNotFound
		}
		return TodoItem{}, err
	}

	return item, nil
}

func saveItem(ctx context.Context, db *sql.DB, item TodoItem) error {
	q := `INSERT ITO todolist(id, title, created_at, done_at) VALUES($1, $2, $3, $4) ON CONFLICT(id) DO UPDATE SET title=$2, done_at=$4`

	_, err := db.ExecContext(ctx, q, item.Id, item.Title, item.CreatedAt, item.DoneAt)

	if err != nil {
		return err
	}

	return nil
}
