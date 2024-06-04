package db

type Debt struct {
	DebtorUsername string
	OwnerUsername  string
	Amount         int
	Currency       string
	Description    string
	Date           string
	Language       string
}
