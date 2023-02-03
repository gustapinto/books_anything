package config

import (
	"fmt"
	"os"
)

const (
	DBDriver = "pgx"
)

var (
	AppAddr    = os.Getenv("APP_ADDR")
	AppSecret  = os.Getenv("APP_SECRET")
	DBHost     = os.Getenv("DB_HOST")
	DBPort     = os.Getenv("DB_PORT")
	DBUser     = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName     = os.Getenv("DB_NAME")

	DBDsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DBHost, DBPort, DBPassword, DBUser, DBName)
)
