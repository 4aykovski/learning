package main

import (
	"flag"
	tgClient "github.com/4aykovski/learning/tree/main/golang/telegram-bot/clients/telegram"
	event_consumer "github.com/4aykovski/learning/tree/main/golang/telegram-bot/consumer/event-consumer"
	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/events/telegram"
	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/storage/files"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
