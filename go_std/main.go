package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/gustapinto/books_rest/go_std/config"
	"github.com/gustapinto/books_rest/go_std/controller"
	"github.com/gustapinto/books_rest/go_std/middleware"
	"github.com/gustapinto/books_rest/go_std/repository"
)

func main() {
	db, err := sql.Open(config.DBDriver, config.DBDsn)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.Default()
	usersRepository := repository.NewUsersRepository(db)
	usersController := controller.NewUsersController(usersRepository)

	if err := repository.AutoMigrate(db, usersRepository); err != nil {
		log.Fatal(err)
	}

	http.Handle("/ping/", middleware.Logging(logger, new(controller.PingController)))
	http.Handle("/user/", middleware.Logging(logger, usersController))

	if err := http.ListenAndServe(config.AppAddr, nil); err != nil {
		log.Fatal(err)
	}
}
