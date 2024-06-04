package db

type Debt struct {
	DebtorUsername string
	OwnerUsername  string
	Amount         int
	Currency       string
	Date           string
	Language       string
}
