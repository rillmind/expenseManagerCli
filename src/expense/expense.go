package expense

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func AddExpense(description string, amount int, db *sql.DB) {
	ld := time.Now().Format("2006-01-02")

	insertExpense := `
		insert into expenses(description, amount, createdAt, updatedAt) values (?, ?, ?, ?)
	`

	res, err := db.Exec(insertExpense, description, amount, ld, ld)

	if err != nil {
		fmt.Println(err)
	}

	id, err := res.LastInsertId()

	fmt.Printf("Expense added! (ID: %d)\n", id)
}

func ListExpenses(db *sql.DB) {
	var id, amount int
	var description, createdAt, updatedAt string

	listExpenses := `
		select id, description, amount, createdAt, updatedAt from expenses
	`
	rows, err := db.Query(listExpenses)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &description, &amount, &createdAt, &updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("ID: %02d, description: %s, amount: %d, createdAt: %v\n", id, description, amount, createdAt)
	}
}
