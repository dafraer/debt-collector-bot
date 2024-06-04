package bot

import (
	"collector/src/db"
	"collector/src/tgspam"
	"fmt"
	"time"
)

func (b *Bot) Run() error {
	//Open telegram
	if err := tgspam.OpenTelegram(); err != nil {
		return err
	}

	for i := 1; i < 7; i++ {
		//If 7 days have passed change account
		if i%b.Interval == 0 && i > 1 {
			err := tgspam.ChangeAccount((i % 4))
			if err != nil {
				return err
			}
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
			return err
		}

		//Send messages to the debtors
		for j := 0; j < len(debtors); j++ {
			if debtors[j].Language == "ru" {
				if err := tgspam.SendMessage(debtors[j].DebtorUsername, fmt.Sprintf(b.MessageFormatRu, debtors[j].DebtorUsername, debtors[j].Amount, debtors[j].Currency, debtors[j].OwnerUsername, debtors[j].OwnerUsername)); err != nil {
					return err
				}
			} else {
				if err := tgspam.SendMessage(debtors[j].DebtorUsername, fmt.Sprintf(b.MessageFormatEn, debtors[j].DebtorUsername, debtors[j].Amount, debtors[j].Currency, debtors[j].OwnerUsername, debtors[j].OwnerUsername)); err != nil {
					return err
				}
			}
			time.Sleep(3 * time.Second)
		}

		//Update dates of last notification
		if err := db.UpdateDebtInfo(); err != nil {
			return err
		}

		//Wait for another day
		<-day
	}
	return nil
}
