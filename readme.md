<h1 align="center"> Expense Manager CLI </h1>

<p align="center">
A quick command line expense manager written in Golang. This application allows you to add, update
delete, list, and mark the status of expenses. Expenses are stored in a sqlite file for persistence.
</p>

<p align="center">Project URL: https://roadmap.sh/projects/expense-tracker</p>

<p align="center">
  <img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/palette/macchiato.png">
</p>

## Requirements

- Go 1.24.8 or higher
- git

## Installation

Linux:

```sh
git clone https://github.com/rillmind/expenseManagerCli.git
cd expenseManagerCli
go run ./cmd/seed
go build -o expense-tracker ./cmd/server
```

Windows: 

```sh
git clone https://github.com/rillmind/expenseManagerCli.git
cd expenseManagerCli
go run ./cmd/seed
go build -o expense-tracker.exe ./cmd/server
```

note: This aplication was build using sqlite, so just will run in this project directory. And for Windows users: you need to install
a C compiler because of the sqlite driver.

## Run example

```sh
./expense-tracker add --description "Lunch" --amount 20.0
# Expense added successfully (ID: 1)

./expense-tracker add --description "Dinner" --amount 10.0
# Expense added successfully (ID: 2)

./expense-tracker list
# ID  Date        Amount Description
# 1   2024-08-06  $20    Lunch
# 2   2024-08-06  $10    Dinner

./expense-tracker summary
# Total expenses: $30

./expense-tracker delete --id 2
# Expense deleted successfully

./expense-tracker summary
# Total expenses: $20

./expense-tracker summary --month 8
# Total expenses for August: $20
```

## Code information

### Functions

- **AddExpense :** Adds a expense with given description and amount
- **ListExpenses :** List all expenses
- **SummaryExpenses :** Gives a summary with a sum of all expenses
- **SummaryExpensesByMonth :** Gives a summary with a sum of all expenses by month
- **UpdateExpense :** Updates a expense
- **DeleteExpense :** Deletes a expense

```go
func AddExpense(description string, amount float64, db *sql.DB)

func ListExpenses(db *sql.DB)

func SummaryExpenses(db *sql.DB)

func SummaryExpensesByMonth(month int, db *sql.DB)

func UpdateExpense(id int, description string, amount float64, db *sql.DB)

func DeleteExpense(id int, db *sql.DB)
```