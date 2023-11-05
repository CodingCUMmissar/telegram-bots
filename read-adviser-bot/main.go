package main

import (
	"clients/telegram"
	"flag"

	"log" //"github.com/rs/zerolog/log"
)

func main() {

	tgClient = telegram.New(mustHost(), mustToken())

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)

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

func mustHost() string {
	host := flag.String(
		"telegram-host",
		"",
		"host for telegram",
	)

	flag.Parse()

	if *host == "" {
		log.Fatal("token is not specified")
	}

	return *host
}
