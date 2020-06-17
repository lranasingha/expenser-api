package data

import (
	"database/sql"
	"expenser-api/service"
	"fmt"
	"log"
)

type DbCredentials struct {
	Username string
	Password string
}
type DbProperties struct {
	Type        string
	Host        string
	Port        int
	Credentials DbCredentials
	DbName      string
}

type ExpenseDbClient interface {
	Insert(expense service.Expense)
}

func MakeDb(dbProperties DbProperties) *sql.DB {
	db, err := sql.Open(dbProperties.Type, buildConnString(dbProperties.Type, dbProperties))
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(2)
	return db
}

func buildConnString(dbType string, dbProperties DbProperties) string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		dbProperties.Type,
		dbProperties.Credentials.Username,
		dbProperties.Credentials.Password,
		dbProperties.Host,
		dbProperties.Port,
		dbProperties.DbName,
	)
}

func (dbProperties DbProperties) Insert(expense service.Expense) {

}
