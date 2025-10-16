package expense

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func AddExpense(description string, amount float64, db *sql.DB) {
	amountInt := int(amount * 100)

	// Essa exata string é a que define o formato YYYY-MM-DD.
	ld := time.Now().Format("2006-01-02")

	query := `
		insert into expenses(description, amount, createdAt, updatedAt) values (?, ?, ?, ?)
	`

	res, err := db.Exec(query, description, amountInt, ld, ld)

	if err != nil {
		fmt.Println(err)
	}

	id, err := res.LastInsertId()

	fmt.Printf("Expense added! (ID: %d)\n", id)
}

func ListExpenses(db *sql.DB) {
	var (
		id, amount           int
		description          string
		createdAt, updatedAt time.Time
	)

	query := `
		select id, description, amount, createdAt, updatedAt
		from expenses
	`
	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err)
	}

	// Define espacos fixos na string para alinhar com a query do BD
	fmt.Printf("%-3s %-11s %-11s %s\n", "ID", "DATE", "AMOUNT", "DESCRIPTION")
	for rows.Next() {
		err := rows.Scan(&id, &description, &amount, &createdAt, &updatedAt)

		idStr := fmt.Sprintf("%2d", id)
		amountF := float64(amount) / 100
		amountS := fmt.Sprintf("%.2f", amountF)

		createdAtTime := createdAt.Format("2006-01-02")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-3s %-11s R$%-9s %s\n", idStr, createdAtTime, amountS, description)
	}
}

func SummaryExpenses(db *sql.DB) {
	var amount int

	query := `
		select sum(amount)
		from expenses
	`

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err)
	}

	rows.Scan(&amount)

	amountF := float64(amount) / 100
	amountS := fmt.Sprintf("%.2f", amountF)

	fmt.Printf("Total expenses: R$%s\n", amountS)
}

func SummaryExpensesByMonth(month int, db *sql.DB) {
	var amount int

	query := `
		select sum(amount)
		from expenses
		where
			strftime('%m', updatedAt) = ?
	`

	monthStr := fmt.Sprintf("%02d", month)

	err := db.QueryRow(query, monthStr).Scan(&amount)

	if err != nil {
		if err == sql.ErrNoRows {
			amount = 0
		} else {
			log.Fatalln("Summary by expense", err)
		}
	}

	amountF := float64(amount) / 100
	amountS := fmt.Sprintf("%.2f", amountF)

	fmt.Printf("Total expenses: %s\n", amountS)
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
	query := `
		delete
		from expenses
		where id = ?
	`

	res, err := db.Exec(query, id)

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
