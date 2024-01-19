package main

import (
	"flag"
	tgClient "github.com/4aykovski/learning/tree/main/golang/telegram-bot/clients/telegram"
	event_consumer "github.com/4aykovski/learning/tree/main/golang/telegram-bot/consumer/event-consumer"
	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/events/telegram"
	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/logger"
	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/storage/files"
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

	logger.Info.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		logger.Err.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")
	flag.Parse()

	if *token == "" {
		logger.Err.Fatal("token is not specified")
	}

	return *token
}
