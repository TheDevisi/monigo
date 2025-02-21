package telegram

import (
	"context"
	"monigo/config"
	"monigo/utils"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	log "github.com/sirupsen/logrus"
)

func SendDiskUsageInfoStatus(b *bot.Bot) {
	ownerID := config.GetOwnerID()
	b.RegisterHandler(bot.HandlerTypeMessageText, "💿", bot.MatchTypeContains, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		log.WithFields(log.Fields{
			"chatID": update.Message.Chat.ID,
			"user":   update.Message.From.Username,
		}).Debug("Disk usage info requested")

		// Checking if the user is the owner
		if update.Message.Chat.ID != ownerID {
			log.Warn("Unauthorized disk usage info request from user:", update.Message.From.Username)
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "🚫 You don't have permission! / У вас нет прав!",
			})
			if err != nil {
				log.Error("Failed to send unauthorized message:", err)
			}
			log.Warn("Unauthorized disk usage info request from user:", update.Message.From.Username)
			return
		}

		// Send initial "collecting" message
		msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "⏳ Collecting disk usage information...",
		})
		if err != nil {
			log.Error("Failed to send initial message:", err)
			return
		}

		// Collecting info about disk usage using utils/get_disk_usage.go
		diskUsage := utils.GetDiskUsage()

		// Edit the message with disk usage information
		_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    update.Message.Chat.ID,
			MessageID: msg.ID,
			Text:      diskUsage,
		})
		if err != nil {
			log.Error("Failed to edit message with disk usage info:", err)
		} else {
			log.Info("Disk usage info sent successfully")
		}
	})
}
