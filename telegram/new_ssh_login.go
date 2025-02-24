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

func SendMessageLogin(b *bot.Bot) {
	ownerID := config.GetOwnerID()

	// Start goroutine for continuous monitoring
	go func() {
		for {
			// If new login detected
			if monitoring.MonitorSshLogins() {
				// Form the message text
				message := "🔐 <b>New SSH login detected!</b>\n\n" +
					"⚠️ Please check your system for unauthorized access."

				// Send message to owner
				_, err := b.SendMessage(context.Background(), &bot.SendMessageParams{
					ChatID:    ownerID,
					Text:      message,
					ParseMode: models.ParseModeHTML,
				})

				if err != nil {
					log.Error("Failed to send SSH login notification:", err)
				} else {
					log.Info("SSH login notification sent successfully")
				}
			}

			// Small pause before next check
			time.Sleep(100 * time.Millisecond)
		}
	}()
}
