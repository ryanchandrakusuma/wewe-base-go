package routes

import (
	"ryanchandrakusuma/wewe-base-go/api/book/delivery/controller"
	"ryanchandrakusuma/wewe-base-go/api/book/repository"
	"ryanchandrakusuma/wewe-base-go/api/book/usecase"
	"ryanchandrakusuma/wewe-base-go/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewBookRouter(env *config.Env, timeout time.Duration, db *pgx.Conn, group *gin.RouterGroup) {
	br := repository.NewBookRepository(db)
	bc := &controller.BookController{
		BUsecase: usecase.NewBookUsecase(br, timeout),
	}
	group.GET("/books", bc.Fetch)
}