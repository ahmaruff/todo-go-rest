package main

import (
	"os"
	"todo-go/todo"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		os.Setenv("TODO_SERVER_ADDRESS", "0.0.0.0")
		os.Setenv("TODO_SERVER_PORT", "80")
	}

	serverAddress := os.Getenv("TODO_SERVER_ADDRESS")
	serverPort := os.Getenv("PORT")

	todo.SetupDatabase()

	e := echo.New()
	todo.InitTodoRoutes(e)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	server := serverAddress + ":" + serverPort
	log.Info().Msg("Starting up server...")

	err = e.Start(server)
	e.Logger.Fatal(err)
}
