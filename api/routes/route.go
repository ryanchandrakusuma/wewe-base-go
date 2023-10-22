package routes

import (
	"database/sql"
	"ryanchandrakusuma/wewe-base-go/config"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *config.Env, timeout time.Duration, db *sql.DB, gin *gin.Engine) {

}