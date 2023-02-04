package main

import (
	"database/sql"
	"log"
	"net/http"
	"regexp"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/gustapinto/books_rest/go_std/config"
	"github.com/gustapinto/books_rest/go_std/controller"
	"github.com/gustapinto/books_rest/go_std/middleware"
	"github.com/gustapinto/books_rest/go_std/model"
	"github.com/gustapinto/books_rest/go_std/repository"
)

func main() {
	logger := log.Default()

	db, err := sql.Open(config.DBDriver, config.DBDsn)
	if err != nil {
		logger.Fatal(err)
	}

	if err := model.AutoMigrate(db, new(model.User), new(model.Author)); err != nil {
		logger.Fatal(err)
	}

	usersRepository := repository.NewUsersRepository(db)
	authorsRepository := repository.NewAuthorsRepository(db)

	RegisterRoutes(logger, map[string]http.Handler{
		"/ping":   controller.NewPingController(),
		"/user":   controller.NewUsersController(usersRepository),
		"/auth":   controller.NewAuthController(usersRepository),
		"/author": middleware.Auth(controller.NewAuthorsController(authorsRepository)),
	})
}

func RegisterRoutes(logger *log.Logger, routeMapping map[string]http.Handler) {
	for route, handler := range routeMapping {
		http.Handle(route, middleware.Logging(logger, handler))

		if matched, _ := regexp.Match(`\/[a-z|A-Z]*\/`, []byte(route)); !matched {
			http.Handle(route+"/", middleware.Logging(logger, handler))
		}
	}

	if err := http.ListenAndServe(config.AppAddr, nil); err != nil {
		logger.Fatal(err)
	}
}
