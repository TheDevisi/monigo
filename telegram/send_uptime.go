package telegram

import (
	"context"
	"monigo/config"
	"monigo/utils"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	log "github.com/sirupsen/logrus"
)

func Send_uptime(b *bot.Bot) {
	ownerID := config.GetOwnerID()
	b.RegisterHandler(bot.HandlerTypeMessageText, "⏱️", bot.MatchTypeContains, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		log.WithFields(log.Fields{
			"chatID": update.Message.Chat.ID,
			"user":   update.Message.From.Username,
		}).Debug("Uptime info requested")

		// Checking if the user is the owner
		if update.Message.Chat.ID != ownerID {
			log.Warn("Unauthorized uptime info request from user:", update.Message.From.Username)
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "🚫 You don't have permission! / У вас нет прав!",
			})
			if err != nil {
				log.Error("Failed to send unauthorized message:", err)
			}
			log.Warn("Unauthorized uptime info request from user:", update.Message.From.Username)
			return
		}

		// Send initial "collecting" message
		msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "⏳ Collecting uptime information...",
		})
		if err != nil {
			log.Error("Failed to send initial message:", err)
			return
		}

		// Collecting uptime info
		uptime := utils.Uptime()

		// Edit the message with uptime information
		_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    update.Message.Chat.ID,
			MessageID: msg.ID,
			Text:      uptime,
		})
		if err != nil {
			log.Error("Failed to edit message with uptime info:", err)
		} else {
			log.Info("Uptime info sent successfully")
		}
	})
}
