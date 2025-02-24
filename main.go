package main

import (
	"context"
	"monigo/config"
	"monigo/telegram"

	"github.com/go-telegram/bot"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Configure logrus
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		DisableColors: false,
	})
}

func main() {
	// Load environment variables from .env
	config.LoadEnv()

	// Set log level
	log.SetLevel(log.InfoLevel)

	log.Info("Starting Monigo bot...")

	opts := []bot.Option{
		bot.WithDefaultHandler(telegram.DefaultHandler),
	}

	log.Info("Environment variables loaded successfully")

	// Importing the token from .env
	BOT_TOKEN := config.GetToken()
	log.Debug("Bot token retrieved")

	b, err := bot.New(BOT_TOKEN, opts...)
	if err != nil {
		log.Fatal("Failed to create bot instance: ", err)
	}
	log.Info("Bot instance created successfully")

	// Register command handlers
	log.Info("Registering command handlers...")
	telegram.Send_uptime(b)
	telegram.SendRamInfoStatus(b)
	telegram.SendDiskUsageInfoStatus(b)
	telegram.SendCpuInfoStatus(b)
	telegram.SendMessageLogin(b)
	log.Info("Command handlers registered successfully")

	// Register start command to show keyboard
	telegram.Start(b)
	log.Info("Start command handler registered")

	log.Info("Bot is now running. Press Ctrl+C to exit")
	b.Start(context.Background())
}
