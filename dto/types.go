package dto

//the TitleCase name is required to correctly decode the JSON
type Expense struct {
	Name        string
	Description string
	Category    string
	Payload     string
}
