package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	usecase := NewUseCase(db)

	p, err := usecase.GetProduct(1)

	if err != nil {
		panic(err)
	}

	fmt.Println(p.Name)
}
