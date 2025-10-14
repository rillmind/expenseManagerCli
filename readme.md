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