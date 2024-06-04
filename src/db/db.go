package db

type Debt struct {
	debtorUsername string
	OwnerUsername  string
	Amount         int
	Description    string
	Date           string
	Language       string
}
