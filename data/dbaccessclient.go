package data

import (
	"database/sql"
	"expenser-api/dto"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DbCredentials struct {
	Username string
	Password string
}
type DbClient struct {
	Type        string
	Host        string
	Port        int
	Credentials DbCredentials
	DbName      string
}

type ExpenseDbClient interface {
	Insert(expense dto.Expense)
}

func (client DbClient) MakeDb() *sql.DB {
	db, err := sql.Open(client.Type, client.buildConnString())
	if err != nil {
		panic(err)
	}
	log.Print("connected to DB")
	db.SetMaxOpenConns(2)
	defer db.Close()
	return db
}

func (client DbClient) buildConnString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		client.Host,
		client.Port,
		client.Credentials.Username,
		client.Credentials.Password,
		client.DbName,
	)
}

const insertExpense = "INSERT INTO expense (category, name, description, payload) VALUES(?,?,?, ?);"

func (client DbClient) Insert(expense dto.Expense, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare(insertExpense)

	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	if r, err := stmt.Exec(expense.UserId, expense.Category, expense.Payload); err != nil {
		return -1, err
	} else {
		id, _ := r.LastInsertId()
		return id, nil
	}
}
