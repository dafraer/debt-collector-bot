package bot

type Bot struct {
	Interval        int
	MessageFormatRu string
	MessageFormatEn string
}

func NewBot(interval int, msgRu, msgEn string) *Bot {
	return &Bot{
		Interval:        interval,
		MessageFormatRu: msgRu,
		MessageFormatEn: msgEn,
	}
}
