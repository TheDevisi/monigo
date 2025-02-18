package main

import (
	"context"
	"log"
	"monigo/config"

	"monigo/telegram"

	"github.com/go-telegram/bot"
	// "github.com/joho/godotenv"
)

func main() {
	config.LoadEnv()
	// Importing the token from .env (see config/config.go)
	BOT_TOKEN := config.GetToken()
	b, err := bot.New(BOT_TOKEN) // Creating a new bot and catching an error (if it exists)
	if err != nil {
		log.Fatal(err)
	}
	telegram.Send_uptime(b)       // Sending uptime. This function using /utils/uptime.go
	telegram.SendRamInfoStatus(b) // Sending RAM info. This function using /utils/ram.go

	b.Start(context.Background()) // Starting the bot
}
