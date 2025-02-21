package telegram

import (
	"context"
	"monigo/config"
	"monigo/utils"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	log "github.com/sirupsen/logrus"
)

// SendRamInfoStatus sends RAM info to the user
func SendCpuInfoStatus(b *bot.Bot) {
	ownerID := config.GetOwnerID()
	b.RegisterHandler(bot.HandlerTypeMessageText, "💻 CPU", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		log.WithFields(log.Fields{
			"chatID": update.Message.Chat.ID,
			"user":   update.Message.From.Username,
		}).Debug("CPU info requested")

		// Checking if the user is the owner
		if update.Message.Chat.ID != ownerID {
			log.Warn("Unauthorized CPU info request from user:", update.Message.From.Username)
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "🚫 You don't have permission! / У вас нет прав!",
			})
			log.Warn("Unauthorized CPU info request from user:", update.Message.From.Username)

			if err != nil {
				log.Error("Failed to send unauthorized message:", err)
			}
			return

		}

		// Collecting info about CPU
		cpuInfo := utils.GetCPU()

		// Sending the answer to the user
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      cpuInfo,
			ParseMode: models.ParseModeHTML,
		})
		if err != nil {
			log.Error("Failed to send CPU info:", err)
		} else {
			log.Info("CPU info sent successfully")
		}
	})
}
