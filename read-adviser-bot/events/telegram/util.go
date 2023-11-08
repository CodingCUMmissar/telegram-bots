package telegram

import (
	"telegram-client"

	"github.com/goware/urlx"
)

func NewMessageSender(ChatID string, tg *telegram.Client) func(text string) error {
	return func(msg string) error {
		return tg.SendMessage(ChatID, msg)
	}
}

func isAddCommand(text string) bool {
	_, err := urlx.NormalizeString(text)
	return err == nil
}

func isURL(text string) bool {
	_, err := urlx.Parse(text)
	return err == nil
}
