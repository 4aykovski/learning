package telegram

import (
	"errors"

	telegram2 "github.com/4aykovski/learning/tree/main/golang/telegram-bot/internal/clients/telegram"
	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/internal/events"
	"github.com/4aykovski/learning/tree/main/golang/telegram-bot/internal/storage"
	error_wrapper "github.com/4aykovski/learning/tree/main/golang/telegram-bot/lib/error-wrapper"
)

type Processor struct {
	tg      *telegram2.Client
	offset  int
	storage storage.Storage
}

type Meta struct {
	ChatID   int
	UserName string
}

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType  = errors.New("unknown meta type")
)

func New(client *telegram2.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      client,
		storage: storage,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, error_wrapper.Wrap("can't get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)
	default:
		return error_wrapper.Wrap("can't process message", ErrUnknownEventType)
	}
}

func (p *Processor) processMessage(event events.Event) (err error) {
	defer func() { err = error_wrapper.WrapIfErr("can't process message", err) }()

	meta, err := meta(event)
	if err != nil {
		return err
	}

	if err := p.doCmd(event.Text, meta.ChatID, meta.UserName); err != nil {
		return err
	}

	return nil
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, error_wrapper.Wrap("can't get meta", ErrUnknownMetaType)
	}

	return res, nil
}

func event(upd telegram2.Update) events.Event {
	updType := fetchType(upd)
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}

	if updType == events.Message {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.ID,
			UserName: upd.Message.From.Username,
		}
	}

	return res
}

func fetchText(upd telegram2.Update) string {
	if upd.Message == nil {
		return ""
	}

	return upd.Message.Text
}

func fetchType(upd telegram2.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}

	return events.Message
}
