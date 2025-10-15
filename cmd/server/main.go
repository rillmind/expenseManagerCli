package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rillmind/expenseManagerCli/src/expense"
	"github.com/spf13/pflag"
)

func main() {
	descriptionFlag := pflag.StringP("description", "d", "", "Description of the expense.")
	amountFlag := pflag.IntP("amount", "a", 0, "Amount of the expense")

	pflag.Parse()

	args := pflag.Args()

	db, err := sql.Open("sqlite3", "./db/sqlite.db")

	if err != nil {
		fmt.Println(err)
	}

	switch args[0] {
	case "add":
		var description string
		var amount int
		var err error

		if len(args) > 3 {
			log.Fatal("Too much arguments!")
		}

		if *descriptionFlag != "" || *amountFlag != 0 {
			description = *descriptionFlag
			amount = *amountFlag
		} else {
			description = args[1]
			amount, err = strconv.Atoi(args[2])

			if err != nil {
				fmt.Println(err)
			}
		}

		expense.AddExpense(description, amount, db)

	case "list":
		if len(args) > 1 {
			log.Fatal("Too much arguments!")
		}

		expense.ListExpenses(db)

	case "summary":
		if len(args) > 2 {
			log.Fatal("Too much arguments!")
		}

		expense.SummaryExpenses(db)

	case "delete":
		if len(args) > 2 {
			log.Fatal("Too much arguments!")
		}

		id, err := strconv.Atoi(args[1])

		if err != nil {
			log.Fatalln(err)
		}

		expense.DeleteExpense(id, db)

	default:
		fmt.Printf("Command not recognized: %v!", args[0])
	}
}
