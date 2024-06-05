package main

import (
	bot "collector/src/bot"
	"log"
)

const (
	DefaultInterval        = 7
	DefaultMessageFormatRu = "Привет, %v! Хочу напомнить, что ты должен %v %v пользователю %v. Было бы здорово, если ты сможешь вернуть деньги в ближайшее время.  Если есть какие-то вопросы или тебе нужна дополнительная информация, напиши %v!  Спасибо и надеюсь на понимание."
	DefaultMessageFormatEn = "Hey %v!  Just a quick reminder that you owe %v %v to the user %v. It would be great if you could return the money soon.  If you have any questions or need more information, just let %v know!  Thanks and hope you understand."
)

func main() {
	tgbot := bot.NewBot(1, DefaultMessageFormatRu, DefaultMessageFormatEn)
	if err := tgbot.Run(); err != nil {
		log.Fatal(err)
	}
}
