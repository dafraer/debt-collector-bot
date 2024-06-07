package bot

import "collector/src/logger"

type Bot struct {
	Logger          logger.Logger
	Interval        int
	MessageFormatRu string
	MessageFormatEn string
}

func NewBot(interval int, msgRu, msgEn string) *Bot {
	return &Bot{
		Logger:          logger.NewLogger(3),
		Interval:        interval,
		MessageFormatRu: msgRu,
		MessageFormatEn: msgEn,
	}
}
