package routes

import (
	"ryanchandrakusuma/wewe-base-go/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Setup(env *config.Env, timeout time.Duration, db *pgx.Conn, gin *gin.Engine) {

	publicRouter := gin.Group("")
	// All Public APIs
	NewBookRouter(env, timeout, db, publicRouter)
	
}