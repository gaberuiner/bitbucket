package events

import "testbot/client/telegram"

type Fetcher interface {
	Fetch(limit int) ([]telegram.Update, error)
}

type Process interface {
	Process(e telegram.Update) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
}
