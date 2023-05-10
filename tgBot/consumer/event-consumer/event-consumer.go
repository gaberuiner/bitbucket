package event_consumer

import (
	"log"
	"testbot/client/telegram"
	"testbot/events"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Process
	batchSize int
}

func New(fetcher events.Fetcher, process events.Process, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: process,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf(err.Error())
			continue
		}
		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}
	}
}

func (c *Consumer) handleUpdates(updates []telegram.Update) error {
	for _, upd := range updates {
		log.Printf("new update %s", upd.Message.Text)
		if err := c.processor.Process(upd); err != nil {
			log.Printf("can't handle update: %s", err.Error())
			continue
		}
	}
	return nil
}
