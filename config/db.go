package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

func NewPqDatabase(env *Env) *pgx.Conn {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
	env.DBHost, 
	env.DBUser, 
	env.DBPass, 
	env.DBName, 
	env.DBPort)
	
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db
}