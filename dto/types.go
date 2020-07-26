package dto

//the TitleCase name is required to correctly decode the JSON
type Expense struct {
	Id          int64
	Name        string
	Description string
	Category    string
	Payload     string
}
