package db

import "time"

const (
	//Костыль от Артема
	Russian = false
	English = true
)

type Debt struct {
	Id             int    `gorm:"primary_key"`
	Currency       string `gorm:"column:currency"`
	Amount         int    `gorm:"column:debt_amount"`
	DebtorUsername string `gorm:"column:debtor_username"`
	OwnerUsername  string `gorm:"column:owner_username"`
	//handicap from Artyom
	Language bool      `gorm:"column:is_english"`
	Time     time.Time `gorm:"column:date_time"`
}
