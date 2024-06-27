package main

import (
	bot "collector/src/bot"
)

func main() {
	tgbot := bot.NewBot(1, bot.DefaultMessageFormatRu, bot.DefaultMessageFormatEn)
	tgbot.Logger.Info("New bot created")
	if err := tgbot.Run(); err != nil {
		tgbot.Logger.Error("Error running bot", "error", err)
	}
}
