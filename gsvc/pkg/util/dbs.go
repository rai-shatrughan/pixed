package util

import (
	"database/sql"
	"fmt"
	_ "time"

	_ "github.com/go-sql-driver/mysql"
)

type QueryStatusCode int8

const (
	SuccessQuery QueryStatusCode = iota
	FailedQuery
)

type DbStore struct {
	db         *sql.DB
	StatusCode int8
	err        error
}

func (dbs *DbStore) New() {
	dbs.db, dbs.err = sql.Open("mysql", "root:srpass@tcp(172.18.0.81:3306)/test")
	if dbs.err != nil {
		panic(dbs.err)
	}
	// See "Important settings" section.
	dbs.db.SetConnMaxLifetime(0)
	dbs.db.SetMaxOpenConns(10)
	dbs.db.SetMaxIdleConns(10)
}

func (dbs *DbStore) Query(query string) QueryStatusCode {
	insert, err := dbs.db.Query(query)
	if err != nil {
		fmt.Println(err)
		return FailedQuery
	}
	defer insert.Close()

	return SuccessQuery
}

func (dbs *DbStore) Close() {
	dbs.db.Close()
}
