package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gustapinto/books_rest/go_gin_sqlx/docs"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/app"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

// @title Books REST - GO + Gin + SQLX
// @version dev
// @description Just a simple book management API written using Go, Gin and SQLX

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and the authorized JWT token
func main() {
	docs.SwaggerInfo.BasePath = "/api"

	db, err := sqlx.Connect("pgx", config.DB_DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	app.RunApiApp(db, router)

	if err := router.Run(config.APP_ADDR); err != nil {
		log.Fatal(err)
	}
}
