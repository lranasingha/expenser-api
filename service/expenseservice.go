package service

import (
	"expenser-api/data"
	"expenser-api/dto"
	"log"
)

func ReadCredentials() data.DbCredentials {
	return data.DbCredentials{
		Username: "root",
		Password: "expmgr-rootpw1",
	}
}

var dbClient = data.DbClient{
	Type:        "postgres",
	Host:        "localhost",
	Port:        5432,
	Credentials: ReadCredentials(),
	DbName:      "expense-db",
}
var db = dbClient.MakeDb()

func AddExpense(expense dto.Expense) {
	log.Printf("received expense - %s", expense)
	var id, err = dbClient.Insert(expense, db)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("inserted expense - %d", id)
	}
}
