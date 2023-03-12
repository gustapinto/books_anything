package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/database/repository"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/transport/rest/handler"
	"github.com/gustapinto/books_rest/go_gin_sqlx/internal/transport/rest/middleware"
	"github.com/jmoiron/sqlx"
	swaggoFiles "github.com/swaggo/files"
	swaggoGin "github.com/swaggo/gin-swagger"
)

func RunApiApp(db *sqlx.DB, router *gin.Engine) {
	userRepository := repository.NewUserRepository(db)
	authorRepository := repository.NewAuthorRepository(db)
	bookRepository := repository.NewBookRepository(db)
	pingController := handler.NewPingController()
	userController := handler.NewUserController(userRepository)
	authController := handler.NewAuthController(userRepository)
	authorController := handler.NewAuthorController(authorRepository)
	bookController := handler.NewBookController(bookRepository)

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
