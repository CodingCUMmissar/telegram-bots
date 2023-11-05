package telegram

import "clients/telegram"

type Processor struct {
	tg *telegram.Client
	offset int
	//storage
}

func New(tg *telegram.Client) *Processor {
	return &Processor{
		tg: tg,
		offset: 0,
	}
}