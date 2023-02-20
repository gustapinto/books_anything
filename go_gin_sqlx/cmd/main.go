package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gustapinto/books_rest/go_gin_sqlx/docs"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/config"
	"github.com/gustapinto/books_rest/go_gin_sqlx/pkg/controller"
	"github.com/gustapinto/books_rest/go_gin_sqlx/pkg/middleware"
	"github.com/gustapinto/books_rest/go_gin_sqlx/pkg/repository"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	swaggoFiles "github.com/swaggo/files"
	swaggoGin "github.com/swaggo/gin-swagger"
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

	userRepository := repository.NewUserRepository(db)
	authorRepository := repository.NewAuthorRepository(db)
	bookRepository := repository.NewBookRepository(db)
	pingController := controller.NewPingController()
	userController := controller.NewUserController(userRepository)
	authController := controller.NewAuthController(userRepository)
	authorController := controller.NewAuthorController(authorRepository)
	bookController := controller.NewBookController(bookRepository)

	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/auth", authController.Login)
		api.GET("/ping", pingController.Pong)
		api.GET("/swagger/*any", swaggoGin.WrapHandler(swaggoFiles.Handler))

		user := api.Group("/user").Use(middleware.Auth)
		{
			user.GET("", userController.All)
			user.GET(":userId", userController.Find)
			user.POST("", userController.Create)
			user.PUT(":userId", userController.Update)
			user.DELETE(":userId", userController.Delete)
		}
		author := api.Group("/author").Use(middleware.Auth)
		{
			author.GET("", authorController.All)
			author.GET(":authorId", authorController.Find)
			author.POST("", authorController.Create)
			author.PUT(":authorId", authorController.Update)
			author.DELETE(":authorId", authorController.Delete)
		}
		book := api.Group("/book").Use(middleware.Auth)
		{
			book.GET("", bookController.All)
			book.GET(":bookId", bookController.Find)
			book.POST("", bookController.Create)
			book.PUT(":bookId", bookController.Update)
			book.DELETE(":bookId", bookController.Delete)
		}
	}

	if err := router.Run(config.APP_ADDR); err != nil {
		log.Fatal(err)
	}
}
