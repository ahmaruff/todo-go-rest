package main

import (
	"todo-go/todo"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db_name := "./todo.db"
	todo.SetDatabase(db_name)
}
