package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectWithDataBase() *sql.DB {
	conexao := "user=hugo dbname=postgres password=sua_senha host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
