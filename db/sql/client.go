package sql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"web-svc/config"

	_ "github.com/lib/pq"
)

func NewClient() *sqlx.DB {
	connect, err := sqlx.Connect("postgres", dsn())
	if err != nil {
		log.Panic(fmt.Errorf("cannot connect db: %v", err))
	}

	return connect
}

func dsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Jakarta", config.DbHost, config.DbPort, config.DbUser, config.DbName, config.DbPassword)
}
