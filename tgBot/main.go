package main

import (
	"flag"
	"log"
	"testbot/client/telegram"
	event_consumer "testbot/consumer/event-consumer"
	tg2 "testbot/events/telegram"
	"testbot/storage/database"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "data/storage.db"
	batchSize   = 100
)

func main() {
	basa, err := database.New(storagePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}
	err = basa.Init()
	if err != nil {
		log.Fatal("can't initialize storage: ", err)
	}
	tgClient := telegram.New(tgBotHost, mustToken())
	eventsProcessor := tg2.New(&tgClient, basa)
	log.Printf("Server strated ...")
	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal(err)
	}
}

func mustToken() string {
	f := flag.String("token-bot", "", "token for telegram assess")
	flag.Parse()
	if *f == "" {
		log.Fatal("no token error")
	}

	return *f
}
