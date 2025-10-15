package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/sqlite.db")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Seed() {
	os.Remove("./db/sqlite.db")

	db := Connect()

	createTable := `
		create table expenses (
			"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"description" text,
			"amount" integer,
			"createdAt" date,
			"updatedAt" date
		)
	`

	statement, err := db.Exec(createTable)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tabela criada!", statement)
}
