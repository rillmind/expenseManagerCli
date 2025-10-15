package expense

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func AddExpense(description string, amount float64, db *sql.DB) {
	amountInt := int(amount * 100)

	ld := time.Now().Format("2006-01-02")

	insertExpense := `
		insert into expenses(description, amount, createdAt, updatedAt) values (?, ?, ?, ?)
	`

	res, err := db.Exec(insertExpense, description, amountInt, ld, ld)

	if err != nil {
		fmt.Println(err)
	}

	id, err := res.LastInsertId()

	fmt.Printf("Expense added! (ID: %d)\n", id)
}

func ListExpenses(db *sql.DB) {
	var id, amount int
	var description string
	var createdAt, updatedAt time.Time

	listExpenses := `
		select id, description, amount, createdAt, updatedAt
		from expenses
	`
	rows, err := db.Query(listExpenses)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%-3s %-11s %-11s %s\n", "ID", "DATE", "AMOUNT", "DESCRIPTION")
	for rows.Next() {
		err := rows.Scan(&id, &description, &amount, &createdAt, &updatedAt)

		idStr := fmt.Sprintf("%2d", id)
		amountFloat := float64(amount) / 100
		amountStr := fmt.Sprintf("%.2f", amountFloat)

		createdAtTime := createdAt.Format("2006-01-02")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-3s %-11s R$%-9s %s\n", idStr, createdAtTime, amountStr, description)
	}
}

func SummaryExpenses(db *sql.DB) {
	var amount int

	summaryExpenses := `
		select sum(amount)
		from expenses
	`

	rows, err := db.Query(summaryExpenses)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		rows.Scan(&amount)

		fmt.Printf("Total expenses: R$%d\n", amount)
	}
}

func UpdateExpense(id int, description string, amount float64, db *sql.DB) {
	amountInt := int(amount * 100)

	updateExpense := `
		update expenses
		set description = ?, amount = ?
		where id = ?
	`

	res, err := db.Exec(updateExpense, description, amountInt, id)

	if err != nil {
		log.Fatal(err)
	}

	l, err := res.RowsAffected()

	if l != 1 {
		log.Fatalf("No rows affected! Expense with ID %v not found.\n", id)
	} else {
		fmt.Printf("Expense with ID %v updated sucessfully!\n", id)
	}
}

func DeleteExpense(id int, db *sql.DB) {
	deleteExpense := `
		delete
		from expenses
		where id = ?
	`

	res, err := db.Exec(deleteExpense, id)

	if err != nil {
		log.Fatalln(err)
	}

	l, err := res.RowsAffected()

	if l != 1 {
		log.Fatalf("No rows affected. ID %v not found\n", id)
	} else {
		fmt.Printf("Expense with ID %v deleted sucessfully\n", id)
	}
}
