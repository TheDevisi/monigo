package telegram

import (
	"context"
	"monigo/utils"

	"monigo/config"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	// log "github.com/sirupsen/logrus"
)

// SendRamInfoStatus sends RAM info to the user
func SendCpuInfoStatus(b *bot.Bot) {
	ownerID := config.GetOwnerID()
	b.RegisterHandler(bot.HandlerTypeMessageText, "/cpu", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		// Checking if the user is the owner
		if update.Message.Chat.ID != ownerID {
			_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "У тебя нет прав!",
			})
			return
		}

		// Collecting info about CPU using utils/get_cpu.go
		cpu := utils.GetCPU()

		// Sending an answer to the user
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   cpu,
		})
	})
}
