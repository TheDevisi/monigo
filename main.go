package main

import (
	"context"
	"log"
	"monigo/config"
	"monigo/telegram"

	"github.com/go-telegram/bot"
)

func main() {
	config.LoadEnv()
	// Importing the token from .env (see config/config.go)
	BOT_TOKEN := config.GetToken()

	b, err := bot.New(BOT_TOKEN) // Creating a new bot and catching an error (if it exists)
	if err != nil {
		log.Fatal(err)
	}

	// Register command handlers
	telegram.Send_uptime(b)
	telegram.SendRamInfoStatus(b)
	telegram.SendDiskUsageInfoStatus(b)
	telegram.SendCpuInfoStatus(b)

	// Register start command to show keyboard
	telegram.Start(b)
	// Starting the bot
	b.Start(context.Background())
}
