package event

import (
	"bot/client"
)

type Fetcher interface {
	Fetch(limit int) ([]client.Update, error)
}

type Processor interface {
	Process(updates []client.Update) error
}
