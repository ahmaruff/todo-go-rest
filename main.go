package main

import (
	"net/http"
	"todo-go/todo"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

func main() {
	todo.SetupDatabase()

	e := echo.New()
	todo.InitTodoRoutes(e)

	log.Info().Msg("Starting up server...")

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
