package main

import (
	"collector/src/tgspam"
	"log"
)

func main() {
	if err := tgspam.OpenTelegram(); err != nil {
		log.Fatal(err)
	}
	/*
		if err := tgspam.SendMessage("@dafraer", "Ты должен мне миллион рублей"); err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)*/
	if err := tgspam.ChangeAccount(2); err != nil {
		log.Fatal(err)
	}
	if err := tgspam.ChangeAccount(3); err != nil {
		log.Fatal(err)
	}
	if err := tgspam.ChangeAccount(1); err != nil {
		log.Fatal(err)
	}

}
