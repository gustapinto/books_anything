package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/app/auth"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/internal/domain/user"
	"github.com/gustapinto/books_rest/go_gin_sqlc_ddd/sqlc/out"
)

func Routes(router gin.IRouter, querier out.Querier) {
	userRepository := &user.UserSqlcRepository{Queries: querier}
	userHandler := &UserHandler{UserRepository: userRepository}

	user := router.Group("/user").Use(auth.AuthMiddleware)
	{
		user.GET("", userHandler.All)
		user.GET(":userId", userHandler.Find)
		user.POST("", userHandler.Create)
		user.PUT(":userId", userHandler.Update)
		user.DELETE(":userId", userHandler.Delete)
	}
}
