package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pyros2097/gromer/example/config"
)

var Query *Queries

func init() {
	sqlDB, err := sql.Open("postgres", config.DATABASE_URL)
	if err != nil {
		panic(err)
	}
	Query = New(sqlDB)
}
