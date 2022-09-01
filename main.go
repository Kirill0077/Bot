package main

import (
	tgClient "First_bot/clients/telegram"
	tgeventconsumer "First_bot/consumer/event-consumer"
	"First_bot/events/telegram"
	"First_bot/storage/mysql"
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
	//sqliteStoragePath = "data/sqlite/storage.db"
	mysqlStoragePath = "root:lyly0_07@tcp(127.0.0.1:3306)/storage"
	batchSize        = 100
)

func main() {
	//s := files.New(storagePath)
	s, err := mysql.New(mysqlStoragePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}
	/*
		if err := s.Init(context.TODO()); err != nil {
			log.Fatal("can't init storage: ", err)
		}
	*/
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := tgeventconsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access telegram bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
