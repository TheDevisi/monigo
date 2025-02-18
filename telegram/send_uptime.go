package telegram

import (
	"context"
	"monigo/utils"

	"monigo/config"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	log "github.com/sirupsen/logrus"
)

func Send_uptime(b *bot.Bot) {

	ownerID := config.GetOwnerID()
	b.RegisterHandler(bot.HandlerTypeMessageText, "/uptime", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message.Chat.ID != ownerID {
			_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "У тебя нет прав!",
			})
			log.Info("Кто-то без прав попытался выполнить команду!") // test new log
			return
		}
		uptime := utils.Uptime()
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   uptime,
		})
	})
}
