package telegram

import (
	"context"
	"monigo/config"
	"monigo/monitoring"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	log "github.com/sirupsen/logrus"
)

func HighCPULoadMessage(b *bot.Bot) {
	ownerID := config.GetOwnerID()
	message := "High CPU load!"
	isNotified := false

	go func() {
		for {
			highLoad := monitoring.MonitorCPU()

			if highLoad && !isNotified {
				_, err := b.SendMessage(context.Background(), &bot.SendMessageParams{
					ChatID:    ownerID,
					Text:      message,
					ParseMode: models.ParseModeHTML,
				})
				if err != nil {
					log.Warn("Cannot send message about high cpu load")
				} else {
					log.Info("Successfully sent a message about high cpu load")
					isNotified = true
				}
			} else if !highLoad && isNotified {
				isNotified = false
				log.Info("CPU load normalized, resetting notification flag")
			}

			time.Sleep(time.Second * 5)
		}
	}()
}
