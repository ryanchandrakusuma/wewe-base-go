package main

import (
	"ryanchandrakusuma/wewe-base-go/config"
	"ryanchandrakusuma/wewe-base-go/routes"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	env := config.NewEnv()

	db := config.NewPqDatabase(env)

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	routes.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}