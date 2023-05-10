package telegram

import (
	"fmt"
	"testbot/client/telegram"
	"testbot/storage/database"
)

type Processor struct {
	tgClient *telegram.Client
	offset   int
	storage  *database.Storage
}

func New(client *telegram.Client, storage *database.Storage) *Processor {
	return &Processor{
		tgClient: client,
		storage:  storage,
	}
}

func (p *Processor) Fetch(limit int) ([]telegram.Update, error) {
	updates, err := p.tgClient.Update(p.offset, limit)
	if err != nil {
		return nil, fmt.Errorf("can't get updates: %w", err)
	}

	p.offset = updates[len(updates)-1].ID + 1

	return updates, nil
}

func (p *Processor) Process(u telegram.Update) error {
	switch u.Message.Text {
	case "":
		return fmt.Errorf("no message recieve")
	default:
		if err := p.doCmd(u.Message.Text, u.Message.Chat.Chat_id, u.Message.From.Username); err != nil {
			return err
		}
		return nil
	}
}
