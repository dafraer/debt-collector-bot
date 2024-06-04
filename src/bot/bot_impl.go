package bot

import (
	"collector/src/db"
	"collector/src/tgspam"
	"fmt"
	"time"
)

func (b *Bot) Run() error {
	for i := 0; i < 1; i++ {
		//Get today's debtors
		debtors, err := db.GetDebtInfo()
		if err != nil {
			return err
		}

		//Send messages to the debtors
		if err := tgspam.OpenTelegram(); err != nil {
			return err
		}
		for i := 0; i < len(debtors); i++ {
			if debtors[i].Language == "ru" {
				if err := tgspam.SendMessage(debtors[i].DebtorUsername, fmt.Sprintf(b.MessageFormatRu, debtors[i].DebtorUsername, debtors[i].Amount, debtors[i].Currency, debtors[i].OwnerUsername, debtors[i].OwnerUsername)); err != nil {
					return err
				}
			} else {
				if err := tgspam.SendMessage(debtors[i].DebtorUsername, fmt.Sprintf(b.MessageFormatEn, debtors[i].DebtorUsername, debtors[i].Amount, debtors[i].Currency, debtors[i].OwnerUsername, debtors[i].OwnerUsername)); err != nil {
					return err
				}
			}
			time.Sleep(5 * time.Second)
		}

		//Update dates of last notification
		if err := db.UpdateDebtInfo(); err != nil {
			return err
		}
	}
	return nil
}
