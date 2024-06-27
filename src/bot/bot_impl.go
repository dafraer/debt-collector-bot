package bot

import (
	"collector/src/db"
	"collector/src/tgspam"
	"fmt"
	"time"
)

func (b *Bot) Run() error {
	currentAccount := 1
	if err := db.Connect(); err != nil {
		return err
	}
	defer func() {
		err := db.Disconnect()
		if err != nil {
			b.Logger.Error("Error when disconnecting from db", "error", err)
		}
	}()
	//Open telegram
	if err := tgspam.OpenTelegram(); err != nil {
		b.Logger.Error("Failed to open telegram", "error", err)
		return err
	}
	b.Logger.Info("Opened telegram")

	for i := 1; ; i++ {
		//If interval amount of days have passed change account
		if i%b.Interval == 0 && i > 1 {
			currentAccount = changeAccountNumber(currentAccount)
			err := tgspam.ChangeAccount(currentAccount)
			if err != nil {
				b.Logger.Error("Error changing account", "error", err)
				return err
			}
			b.Logger.Info("Changed account", "account", currentAccount)
		}

		//launch the day
		//TODO change to day for now its a 1 minute interval to test
		day := make(chan struct{})
		go func() {
			//time.Sleep(24 * time.Hour)
			time.Sleep(1 * time.Minute)
			day <- struct{}{}
			close(day)
		}()

		//Get today's debtors
		debtors, err := db.GetDebtInfo()
		if err != nil {
			b.Logger.Error("Error getting debtors", "error", err)
			return err
		}
		b.Logger.Info("Successfully got debtors", "debtors amount", len(debtors))

		//Send messages to the debtors
		for j := 0; j < len(debtors); j++ {
			if debtors[j].Language == db.Russian {
				if err := tgspam.SendMessage(debtors[j].DebtorUsername, fmt.Sprintf(b.MessageFormatRu, debtors[j].DebtorUsername, debtors[j].Amount, debtors[j].Currency, debtors[j].OwnerUsername, debtors[j].OwnerUsername)); err != nil {
					b.Logger.Error("Error sending message", "error", err, "debtor", debtors[j].DebtorUsername, "language", debtors[j].Language)
					return err
				}
				b.Logger.Info("Message sent successfully", "debtor", debtors[j].DebtorUsername, "language", debtors[j].Language)
				//for some reason messages in cyrillic freeze the bot so more tim for processing is needed
				time.Sleep(30 * time.Second)
			} else {
				if err := tgspam.SendMessage(debtors[j].DebtorUsername, fmt.Sprintf(b.MessageFormatEn, debtors[j].DebtorUsername, debtors[j].Amount, debtors[j].Currency, debtors[j].OwnerUsername, debtors[j].OwnerUsername)); err != nil {
					b.Logger.Error("Error sending message", "error", err, "debtor", debtors[j].DebtorUsername, "language", debtors[j].Language)
					return err
				}
				b.Logger.Info("Message sent successfully", "debtor", debtors[j].DebtorUsername, "language", debtors[j].Language)
			}
			time.Sleep(3 * time.Second)
		}

		//Update dates of last notification
		if err := db.UpdateDebtInfo(debtors); err != nil {
			b.Logger.Error("Error updating debt information", "error", err)
			return err
		}
		b.Logger.Info("Successfully updated debt information")
		//Wait for another day
		<-day
	}
}

func changeAccountNumber(accountNumber int) int {
	if accountNumber >= 3 {
		return 1
	} else {
		return accountNumber + 1
	}
}
