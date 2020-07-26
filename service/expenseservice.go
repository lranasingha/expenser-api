package service

import (
	"expenser-api/data"
	"expenser-api/dto"
	"log"
	"os"
)

func ReadCredentials() data.DbCredentials {
	return data.DbCredentials{
		Username: "expsvcuser",
		Password: os.Getenv("EXP_DB_USR_PW"),
	}
}

var dbClient = data.DbClient{
	Type:        "postgres",
	Host:        "localhost",
	Port:        5432,
	Credentials: ReadCredentials(),
	DbName:      "expense_db",
}
var db = dbClient.MakeDb()

const insertExpense = "INSERT INTO expense_schema.user_expense (name, description, category, payload) VALUES($1,$2,$3,$4);"

func AddExpense(expense dto.Expense) {
	log.Printf("received expense - %s", expense)
	var id, err = dbClient.Insert(insertExpense, expense, db)
	if err != nil {
		log.Panic(err)
	} else {
		log.Printf("inserted expense - %d", id)
	}
}

const selectAll = "SELECT user_id, category, name, description, payload FROM expense_schema.user_expense;"

func GetExpenses() []dto.Expense {
	log.Printf("fetching expenses")
	return dbClient.SelectAll(selectAll, db)
}

const updateExpense = "UPDATE expense_schema.user_expense SET name=$1, category=$2, description=$3,payload=$4 WHERE user_id=$5"

func UpdateExpense(expense dto.Expense) {
	log.Printf("updating expense - %s", expense)
	dbClient.Update(updateExpense, expense)
}
func CleanUp() {
	log.Print("Closing DB connection.")
	if err := db.Close(); err != nil {
		log.Panic("failed to close the DB connection. Quitting.")
	}
}
