package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "macbook"
	password = ""
	dbname   = "orders"
)

func CreateConnection() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError("connection err", err)
	err = db.Ping()

	if err != nil {
		panic(err)
	}
	return db
}

func CheckError(from string, err error) {
	if err != nil {
		fmt.Println("from panit: ", from)
		panic(err)
	}
}
