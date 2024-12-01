package env

import (
	"fmt"
	"log"
	"music-library/internal/config"
	"os"
	"strconv"
)

const (
	PG_HOST_ENV          = "PG_HOST"
	PG_PORT_ENV          = "PG_PORT"
	PG_USER_ENV          = "PG_USER"
	PG_PASSWORD_ENV      = "PG_PASSWORD"
	PG_DATABASE_NAME_ENV = "PG_DATABASE_NAME"
)

type configPG struct {
	host     string
	port     int64
	user     string
	password string
	dbName   string
}

func NewConfigPG() config.ConfigPG {
	host := os.Getenv(PG_HOST_ENV)
	portStr := os.Getenv(PG_PORT_ENV)
	user := os.Getenv(PG_USER_ENV)
	pass := os.Getenv(PG_PASSWORD_ENV)
	dbName := os.Getenv(PG_DATABASE_NAME_ENV)
	if len(portStr) == 0 || len(host) == 0 || len(user) == 0 || len(pass) == 0 {
		log.Fatal("Error parse pg config")
	}

	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		log.Fatal("Error parse pg port to int64")
	}

	return &configPG{
		host:     host,
		port:     port,
		user:     user,
		password: pass,
		dbName:   dbName,
	}
}

func (c configPG) DSN() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", c.host, c.port, c.dbName, c.user, c.password)
}
