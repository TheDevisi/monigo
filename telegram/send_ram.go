package telegram

import (
	"context"
	"monigo/config"
	"monigo/utils"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	log "github.com/sirupsen/logrus"
)

func SendRamInfoStatus(b *bot.Bot) {
	ownerID := config.GetOwnerID()
	b.RegisterHandler(bot.HandlerTypeMessageText, "💾 RAM", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		log.WithFields(log.Fields{
			"chatID": update.Message.Chat.ID,
			"user":   update.Message.From.Username,
		}).Debug("RAM info requested")

		// Checking if the user is the owner
		if update.Message.Chat.ID != ownerID {
			log.Warn("Unauthorized RAM info request from user:", update.Message.From.Username)
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "🚫 You don't have permission! / У вас нет прав!",
			})
			log.Warn("Unauthorized RAM info request from user:", update.Message.From.Username)

			if err != nil {
				log.Error("Failed to send unauthorized message:", err)
			}
			return

		}

		// Send initial "collecting" message
		msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "⏳ Collecting RAM information...",
		})
		if err != nil {
			log.Error("Failed to send initial message:", err)
			return
		}

		// Collecting info about ram using utils/get_memory.go
		memory := utils.GetMemory()

		// Edit the message with RAM information
		_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    update.Message.Chat.ID,
			MessageID: msg.ID,
			Text:      memory,
		})
		if err != nil {
			log.Error("Failed to edit message with RAM info:", err)
		} else {
			log.Info("RAM usage info sent successfully")
		}
	})
}
