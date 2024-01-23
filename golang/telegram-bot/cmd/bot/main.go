package main

import (
	"context"
	"flag"
	"log"

	tgClient "github.com/4aykovski/learning/tree/main/golang/telegram-bot/internal/clients/telegram"
	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/internal/consumer/event-consumer"
	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/internal/events/telegram"
	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/internal/storage/sqlite"
	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/pkg/logger"
)

const (
	tgBotHost         = "api.telegram.org"
	storagePath       = "files_storage"
	sqliteStoragePath = "internal/data/sqlite/storage.db"
	batchSize         = 100
)

func main() {

	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatalf("can't connect to storage: %s", err.Error())
	}

	if err = s.Init(context.TODO()); err != nil {
		log.Fatalf("can't init a storage: %s", err.Error())
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
	)

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	logger.Info.Print("service started")

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
