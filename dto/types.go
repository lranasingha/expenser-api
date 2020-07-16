package dto

//the TitleCase name is required to correctly decode the JSON
type Expense struct {
	UserId   string
	Category string
	Payload  string
}
