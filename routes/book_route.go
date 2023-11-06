package routes

import (
	"ryanchandrakusuma/wewe-base-go/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func NewBookRouter(env *config.Env, timeout time.Duration, db *pgx.Conn, gin *gin.RouterGroup) {
	// ur := repository.NewUserRepository(db, domain.CollectionUser)
	// sc := controller.SignupController{
	// 	SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
	// 	Env:           env,
	// }
	// group.POST("/signup", sc.Signup)
}