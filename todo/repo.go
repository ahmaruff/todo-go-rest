package todo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

var ErrTodoNotFound error = errors.New("todo: not found")

func findItemById(ctx context.Context, tx *sql.Tx, id ulid.ULID) (item TodoItem, err error) {
	q := `SELECT id, title, created_at, done_at FROM todolist WHERE id = ?`

	stmt, err := DB.PrepareContext(ctx, q)

	if err != nil {
		log.Debug().Err(err)
		return TodoItem{}, err
	}

	sqlErr := stmt.QueryRow(id).Scan(&item.Id, &item.Title, &item.CreatedAt, &item.DoneAt)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return TodoItem{}, ErrTodoNotFound
		}
		return TodoItem{}, sqlErr
	}

	return item, nil
}

func saveItem(ctx context.Context, tx *sql.Tx, item TodoItem) error {
	q := `INSERT INTO todolist(id, title, created_at, done_at) VALUES(?, ?, ?, ?) ON CONFLICT (id) DO UPDATE SET title= ?, done_at= ?`

	stmt, err := tx.PrepareContext(ctx, q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&item.Id, &item.Title, &item.CreatedAt, &item.DoneAt, &item.Title, &item.DoneAt)

	if err != nil {
		return err
	}

	return nil
}

func deleteItemById(ctx context.Context, tx *sql.Tx, id ulid.ULID) (bool, error) {
	q := `DELETE FROM todolist WHERE id = ?`

	stmt, err := tx.PrepareContext(ctx, q)

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return false, err
	}

	return true, nil
}
