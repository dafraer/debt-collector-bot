package bot

import "collector/src/logger"

const (
	DefaultInterval        = 7
	DefaultMessageFormatRu = "Привет, %v! Хочу напомнить, что ты должен %v %v пользователю %v. Было бы здорово, если ты сможешь вернуть деньги в ближайшее время.  Если есть какие-то вопросы или тебе нужна дополнительная информация, напиши %v!  Спасибо и надеюсь на понимание."
	DefaultMessageFormatEn = "Hey %v!  Just a quick reminder that you owe %v %v to the user %v. It would be great if you could return the money soon.  If you have any questions or need more information, just let %v know!  Thanks and hope you understand."
)

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
