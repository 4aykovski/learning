package event_consumer

import (
	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/events"
	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/logger"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c *Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			logger.Err.Printf("[ERR] consumer: %s", err.Error())
			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			logger.Err.Print(err)
			continue
		}
	}
}

// потеря событий : ретраи, возвращение в хранилище, фоллбэк, подтверждение для фетчера
// обработка всей пачки : останавливаться при определенном количестве ошибок
// параллельная обработка :
func (c *Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		logger.Info.Printf("got new event: %s", event.Text)

		if err := c.processor.Process(event); err != nil {
			logger.Err.Printf("can't handle event: %s", err.Error())
			continue
		}
	}
	return nil
}
