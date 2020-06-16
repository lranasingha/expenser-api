package service

import "fmt"

func AddExpense(expense Expense) {
	fmt.Printf("received expense - %s", expense)
}
