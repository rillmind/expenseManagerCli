package expense

import (
	"database/sql"
	"fmt"
	"time"
)

func AddExpense(description string, amount int, db *sql.DB) {
	ld := time.Now().Local().String()

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
