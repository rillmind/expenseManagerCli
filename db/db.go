package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Seed() {
	dbFile := "./db/sqlite.db"

	os.Remove(dbFile)

	db, err := sql.Open("sqlite3", dbFile)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

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
