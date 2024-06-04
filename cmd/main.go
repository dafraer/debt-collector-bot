package main

import (
	"collector/src/tgspam"
	"log"
)

func main() {
	if err := tgspam.SendMessage("@fiodop", "Ты должен мне миллион рублей"); err != nil {
		log.Fatal(err)
	}

}
