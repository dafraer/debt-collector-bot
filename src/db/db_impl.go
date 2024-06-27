package db

import (
	"collector/src/bot"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var err error

func GetDebtInfo() ([]Debt, error) {
	temp := make([]Debt, 0)

	db.Table("debts").Select("debts.Id, debts.currency, debts.debt_amount, debts.username as debtor_username, debt_owners.username as owner_username, debt_owners.is_english, debts.date_time").Joins("full join debt_owners on debts.owner_id = debt_owners.id").Where("notify=true ").Scan(&temp)
	debts := make([]Debt, 0)

	for _, debt := range temp {
		if time.Now().After(debt.Time) {
			debts = append(debts, debt)
		}
	}
	return debts, nil
}

func UpdateDebtInfo(debts []Debt) error {
	for _, debt := range debts {
		debt.Time = time.Now().Add(24 * time.Hour * bot.DefaultInterval)
		db.Save(&debt)
	}
	return nil
}
func Connect() error {
	dsn := "host=localhost user=postgres password=qwerty dbname=collectormoneybot port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		return err
	}
	return nil
}

func Disconnect() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// Close the connection
	if err := sqlDB.Close(); err != nil {
		return err
	}
	return nil
}
