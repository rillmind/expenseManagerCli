package main

import (
	"database/sql"
	"fmt"
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

	default:
		fmt.Printf("Command not recognized: %v!", args[0])
	}
}
