package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/api"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/app"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// @title Books REST - GO + Gin + SQLX
// @version dev
// @description Just a simple book management API written using Go, Gin and SQLX

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and the authorized JWT token
func main() {
	api.SwaggerInfo.BasePath = "/api"

	router := gin.Default()

	app.RegisterRoutes(router)

	if err := router.Run(config.APP_ADDR); err != nil {
		log.Fatal(err)
	}
}
