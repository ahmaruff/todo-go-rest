package todo

import (
	"context"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
)

func listItems(ctx context.Context) (TodoList, error) {
	tx, err := DB.Begin()

	if err != nil {
		return TodoList{}, err
	}

	list, err := findAllItems(ctx, tx)

	if err != nil {
		return TodoList{}, err
	}

	tx.Commit()

	return list, nil
}

func createItem(ctx context.Context, title string) (id ulid.ULID, err error) {
	todoItem, err := MakeNewItem(title)

	if err != nil {
		return
	}

	tx, err := DB.Begin()

	if err != nil {
		return
	}

	err = saveItem(ctx, tx, todoItem)

	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()

	if err != nil {
		return
	}

	return todoItem.Id, nil
}

func findItem(ctx context.Context, id ulid.ULID) (item TodoItem, err error) {
	tx, err := DB.Begin()

	if err != nil {
		return
	}

	item, err = findItemById(ctx, tx, id)

	if err != nil {
		return TodoItem{}, err
	}

	err = tx.Commit()

	if err != nil {
		return TodoItem{}, err
	}

	return item, nil
}

func makeItemDone(ctx context.Context, id ulid.ULID) error {
	tx, err := DB.Begin()

	if err != nil {
		return err
	}

	item, err := findItemById(ctx, tx, id)

	if err != nil {
		tx.Rollback()
		return err
	}

	if err = item.MakeDone(); err != nil {
		tx.Rollback()
		return err
	}

	if err = saveItem(ctx, tx, item); err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil

}

func deleteItem(ctx context.Context, id ulid.ULID) error {
	tx, err := DB.Begin()

	if err != nil {
		return err
	}

	_, err = deleteItemById(ctx, tx, id)

	if err != nil {
		tx.Rollback()
		log.Debug().Err(err).Msg(err.Error())
		return err
	}
	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}
