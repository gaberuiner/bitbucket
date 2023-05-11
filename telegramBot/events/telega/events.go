package events

import (
	"bot/client"
	database "bot/database/storage"
	"bot/errors"
)

var (
	limit         = 100
	ErrorMessages = "can't recieve updates"
	ErrorProcess  = "can't process messages"
)

type Processor struct {
	tg      client.Client
	offset  int
	storage *database.Storage
}

type Message struct {
	Message string
	ChatID  int
}

func New(tg client.Client, storage *database.Storage) *Processor {
	return &Processor{
		tg:      tg,
		storage: storage,
	}
}

func (p *Processor) Fetch(limit int) ([]client.Update, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, errors.Error(ErrorMessages, err)
	}
	if len(updates) == 0 {
		return nil, nil
	}
	var messages []client.Update
	for _, upd := range updates {
		if upd.Message.Text == "" {
			continue
		}

		messages = append(messages, upd)
	}
	p.offset = updates[len(updates)-1].ID + 1
	return updates, nil
}

func (p *Processor) Process(updates []client.Update) error {
	for _, upd := range updates {
		if upd.CallbackQuery.Data != "" {

			if err := p.InlineButtonProcess(upd.Message.From.Username, upd.CallbackQuery.From.UserID, upd.CallbackQuery.Data, "Data Types"); err != nil {
				return err
			}
			continue
		}
		if err := p.doCmd(upd.Message.From.Username, upd.Message.From.UserID, upd.Message.Text); err != nil {
			return errors.Error(ErrorProcess, err)
		}
	}
	return nil
}
