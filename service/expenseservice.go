package service

import (
	"expenser-api/data"
	"log"
)

func ReadCredentials() data.DbCredentials {
	return data.DbCredentials{
		Username: "user",
		Password: "pass",
	}
}

var db = data.MakeDb(data.DbProperties{
	Type:        "postgres",
	Host:        "localhost",
	Port:        5432,
	Credentials: ReadCredentials(),
	DbName:      "expense-db",
})

const insertExpense = "INSERT INTO expense (Category, Payload) VALUES(?,?);"

func AddExpense(expense Expense) {
	log.Printf("received expense - %s", expense)
	stmt, err := db.Prepare(insertExpense)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer stmt.Close()

	if r, err := stmt.Exec(expense.Category, expense.Payload); err != nil {
		log.Fatal(err)
		return
	} else {
		id, _ := r.LastInsertId()
		log.Printf("inserted expense %d", id)
	}
}
