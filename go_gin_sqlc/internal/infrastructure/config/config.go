package config

import (
	"fmt"
	"os"
)

var (
	APP_ADDR   = os.Getenv("APP_ADDR")
	APP_SECRET = os.Getenv("APP_SECRET")

	DB_HOST     = os.Getenv("DB_HOST")
	DB_PORT     = os.Getenv("DB_PORT")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME     = os.Getenv("DB_NAME")
	DB_DSN      = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
)
