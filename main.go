package main

import (
	"fdcteam-bot/bot"
	"fdcteam-bot/config"
	"fdcteam-bot/database"
	log "github.com/sirupsen/logrus"
)

func main() {
	setupLogger()

	dbErr := database.Connect()

	if dbErr != nil {
		panic(dbErr)
	}

	botErr := bot.StartBot()

	if botErr != nil {
		panic(botErr)
	}

	<-make(chan struct{})
}

func setupLogger() {
	parsedLevel, err := log.ParseLevel(config.LogLevel)

	if err == nil {
		log.Infof("Log level setado para %s", config.LogLevel)
		log.SetLevel(parsedLevel)
	}
}