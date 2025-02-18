package telegram

import (
	"context"
	"monigo/utils"

	"monigo/config"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	// log "github.com/sirupsen/logrus"
)

func SendDiskUsageInfoStatus(b *bot.Bot) {
	ownerID := config.GetOwnerID()
	b.RegisterHandler(bot.HandlerTypeMessageText, "💿 Диск", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		// Checking if the user is the owner
		if update.Message.Chat.ID != ownerID {
			_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "У тебя нет прав!",
			})
			return
		}

		// Collecting info about disk usage using utils/get_disk_usage.go
		diskUsage := utils.GetDiskUsage()

		// Sending an answer to the user
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   diskUsage,
		})
	})

}
