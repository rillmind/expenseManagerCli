package main

import (
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rillmind/expenseManagerCli/db"
	"github.com/rillmind/expenseManagerCli/src/expense"
	"github.com/spf13/pflag"
)

func main() {
	db := db.Connect()

	dFlag := pflag.StringP("description", "d", "", "Description of the expense.")
	aFlag := pflag.Float64P("amount", "a", 0, "Amount of the expense")
	mFlag := pflag.IntP("month", "m", 0, "Month of the year to restrict")

	pflag.Usage = printUsage

	pflag.Parse()

	args := pflag.Args()

	switch args[0] {
	case "add":
		if len(args) > 3 {
			log.Fatalln("Too much arguments!")
		}

		if len(args) < 3 {
			log.Fatalln("Not enough arguments!")
		}

		description := args[1]
		amount, err := strconv.ParseFloat(args[2], 64)

		if err != nil {
			fmt.Println(err)
		}

		if *dFlag != "" || *aFlag != 0 {
			description = *dFlag
			amount = *aFlag
		}

		expense.AddExpense(description, amount, db)

	case "list":
		if len(args) > 1 {
			log.Fatalln("Too much arguments!")
		}

		if len(args) < 1 {
			log.Fatalln("Not enough arguments!")
		}

		expense.ListExpenses(db)

	case "summary":
		var month int

		if *mFlag != 0 {
			month = *mFlag
			expense.SummaryExpensesByMonth(month, db)
		} else {
			expense.SummaryExpenses(db)
		}

	case "update":
		if len(args) > 4 {
			log.Fatalln("Too much arguments!")
		}

		if len(args) < 4 {
			log.Fatalln("Not enough arguments!")
		}

		id, err := strconv.Atoi(args[1])
		description := args[2]
		amount, err := strconv.ParseFloat(args[3], 64)

		if err != nil {
			log.Fatal(err)
		}

		if *dFlag != "" || *aFlag != 0 {
			description = *dFlag
			amount = *aFlag
		}

		expense.UpdateExpense(id, description, amount, db)

	case "delete":
		if len(args) > 2 {
			log.Fatal("Too much arguments!")
		}

		if len(args) < 2 {
			log.Fatalln("Not enough arguments!")
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

func printUsage() {
	fmt.Printf(`Expense Manager CLI - Manage your expenses from the command line

Usage:
	expense-tracker [options] <command> [arguments]

Options:
	-v    Activate verbose mode, displaying more operation details

Commands:
	add <description> <amount>             Add a new expense
	list                                   List all tasks
	summary                                Show how much you spent at all
	update <id> <description> <amount>     Update a expense
	delete <id>                            Delete a expense

Examples:
	expense-tracker add "Groceries" 12.50
	expense-tracker list
	expense-tracker update 1 "Summer groceries" 12.50
`)
}
