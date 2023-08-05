package config

import (
	"os"
	"strings"
)

const (
	DB_DRIVER = "pgx"
)

var (
	APP_ADDR   = os.Getenv("APP_ADDR")
	APP_SECRET = os.Getenv("APP_SECRET")

	CASSANDRA_HOSTS = strings.Split(os.Getenv("CASSANDRA_HOSTS"), ",")
)
