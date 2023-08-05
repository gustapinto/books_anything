package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/database/repository"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/database/services"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/transport/rest/handler"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/transport/rest/middleware"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/sqlc/out"
	swaggoFiles "github.com/swaggo/files"
	swaggoGin "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine, queries *out.Queries) {
	userRepository := &repository.UserSqlcRepository{Queries: queries}
	authorRepository := &repository.AuthorSqlcRepository{Queries: queries}
	bookRepository := &repository.BookSqlcRepository{}

	userService := &services.UserService{UserRepository: userRepository}
	authorService := &services.AuthorService{AuthorRepository: authorRepository}
	bookService := &services.BookService{BookRepository: bookRepository}

	pingController := &handler.PingController{}
	userController := &handler.UserController{UserService: userService}
	authController := &handler.AuthController{UserService: userService}
	authorController := &handler.AuthorController{AuthorService: authorService}
	bookController := &handler.BookController{BookService: bookService}

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
}
