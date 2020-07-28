package data

import (
	"database/sql"
	"encoding/base64"
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

func (client DbClient) Insert(query string, expense dto.Expense, db *sql.DB) (int64, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Print("recovered from error, returning")
		}
	}()

	imageBytes, decodeErr := base64.StdEncoding.DecodeString(expense.Payload)
	if decodeErr != nil {
		client.Log(decodeErr)
		return -1, decodeErr
	}

	if tx, txErr := db.Begin(); txErr == nil {
		stmt, err := tx.Prepare(query)

		if err != nil {
			client.Rollback(tx)
			return -1, nil
		}

		if r, err := stmt.Exec(expense.Name, expense.Description, expense.Category, imageBytes); err != nil {
			client.Rollback(tx)
			return -1, err
		} else {
			id, _ := r.LastInsertId()
			err := tx.Commit()
			client.Log(err)
			return id, nil
		}
	} else {
		client.Log(txErr)
		return -1, txErr
	}
}

func (client DbClient) Rollback(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		log.Panic(err)
	}
}

func (client DbClient) SelectAll(query string, db *sql.DB) []dto.Expense {
	var expenses []dto.Expense
	if rows, qErr := db.Query(query); qErr != nil {
		client.Log(qErr)
	} else {
		for rows.Next() {
			var expense dto.Expense
			if sErr := rows.Scan(&expense.Id, &expense.Name, &expense.Description, &expense.Category, &expense.Payload); sErr != nil {
				client.Log(sErr)
			}

			expenses = append(expenses, expense)
		}
	}
	return expenses
}

func (client DbClient) Update(query string, expense dto.Expense, db *sql.DB) (int64, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Print("recovered from error, returning")
		}
	}()

	imageBytes, decodeErr := base64.StdEncoding.DecodeString(expense.Payload)
	if decodeErr != nil {
		client.Log(decodeErr)
		return -1, decodeErr
	}

	if tx, txErr := db.Begin(); txErr == nil {
		stmt, err := tx.Prepare(query)

		if err != nil {
			client.Rollback(tx)
			return -1, nil
		}

		if r, exErr := stmt.Exec(expense.Name, expense.Category, expense.Description, imageBytes, expense.Id); exErr != nil {
			client.Rollback(tx)
			return -1, exErr
		} else {
			id, _ := r.LastInsertId()
			txCommitErr := tx.Commit()
			client.Log(txCommitErr)
			return id, nil
		}
	} else {
		client.Log(txErr)
		return -1, txErr
	}
}

func (client DbClient) Log(err error) {
	log.Print(err)
}
