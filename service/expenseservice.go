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

func AddExpense(expense dto.Expense) {
	log.Printf("received expense - %s", expense)
	var id, err = dbClient.Insert(expense, db)
	if err != nil {
		log.Panic(err)
	} else {
		log.Printf("inserted expense - %d", id)
	}
}

func CleanUp() {
	log.Print("cleaning up the service. closing DB connection.")
	if err := db.Close(); err != nil {
		log.Panic("failed to close the DB connection. Quitting.")
	}
}
