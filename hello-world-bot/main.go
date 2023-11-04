package main

import (
	"flag"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Get the Telegram bot token from flag
	bot, err := tgbotapi.NewBotAPI(mustToken())
	tgbotapi.
	if err != nil {
		log.Fatal("Failed to initialize bot:", err)
	}

	// Set up updates channel to receive incoming messages
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal("Failed to get updates channel:", err)
	}

	// Handle incoming updates
	for update := range updates {
		// Verify if the update is a message
		if update.Message == nil {
			continue
		}

		// Create a new message config
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello world!")

		// Send the message
		_, err = bot.Send(msg)
		if err != nil {
			log.Println("Failed to send message:", err)
		}
	}
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
