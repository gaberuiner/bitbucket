package start

import (
	event "bot/events"
	"fmt"
	"time"
)

type Consumer struct {
	fetcher   event.Fetcher
	processor event.Processor
	batchSize int
}

func New(fetcher event.Fetcher, processor event.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c *Consumer) Start() error {
	for {
		Messages, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		if len(Messages) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}
		err = c.processor.Process(Messages)
		if err != nil {
			fmt.Println(err)
			continue
		}

	}
}
