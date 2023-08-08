package app

// func RegisterRoutes(router *gin.Engine, queries *out.Queries) {
// 	userRepository := &repository.UserSqlcRepository{Queries: queries}
// 	authorRepository := &repository.AuthorSqlcRepository{Queries: queries}
// 	bookRepository := &repository.BookSqlcRepository{}

// 	userService := &services.UserService{UserRepository: userRepository}
// 	authorService := &services.AuthorService{AuthorRepository: authorRepository}
// 	bookService := &services.BookService{BookRepository: bookRepository}

// 	pingController := &handler.PingController{}
// 	userController := &handler.UserController{UserService: userService}
// 	authController := &handler.AuthController{UserService: userService}
// 	authorController := &handler.AuthorController{AuthorService: authorService}
// 	bookController := &handler.BookController{BookService: bookService}

// 	api := router.Group("/api")
// 	{
// 		api.POST("/auth", authController.Login)
// 		api.GET("/ping", pingController.Pong)
// 		api.GET("/swagger/*any", swaggoGin.WrapHandler(swaggoFiles.Handler))

// 		author := api.Group("/author").Use(middleware.Auth)
// 		{
// 			author.GET("", authorController.All)
// 			author.GET(":authorId", authorController.Find)
// 			author.POST("", authorController.Create)
// 			author.PUT(":authorId", authorController.Update)
// 			author.DELETE(":authorId", authorController.Delete)
// 		}
// 		book := api.Group("/book").Use(middleware.Auth)
// 		{
// 			book.GET("", bookController.All)
// 			book.GET(":bookId", bookController.Find)
// 			book.POST("", bookController.Create)
// 			book.PUT(":bookId", bookController.Update)
// 			book.DELETE(":bookId", bookController.Delete)
// 		}
// 	}
// }
