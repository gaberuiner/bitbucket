package main

import (
	"bot/client"
	"bot/consumer/start"
	database "bot/database/storage"
	events "bot/events/telega"
	"log"
)

var (
	tgBotHost = "api.telegram.org"
	Token     = "5781780785:AAE5Tc7r5YQXukTKMYPnyBTEV23PFr4t7vs"
	batchSize = 100
)
var storagePath = "data/database.db"

func main() {
	basa, err := database.New(storagePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}
	err = basa.Init()
	if err != nil {
		log.Fatal("can't initialize storage: ", err)
	}
	tgClient := client.New(tgBotHost, Token)
	StartBot := events.New(tgClient, basa)
	log.Printf("Server started ...")
	consumer := start.New(StartBot, StartBot, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal(err)
	}
}
