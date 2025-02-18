package telegram

import (
	"context"
	"monigo/config"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var UserLanguage = make(map[int64]string)

func Start(b *bot.Bot) {
	// Start command handler
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		ownerID := config.GetOwnerID()
		if update.Message.Chat.ID != ownerID {
			_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "🚫 You don't have permission! / У вас нет прав!",
			})
			return
		}

		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "👋 Please select your language / Пожалуйста, выберите язык:",
			ReplyMarkup: CreateLanguageKeyboard(),
		})
	})

	// Language selection handler
	b.RegisterHandler(bot.HandlerTypeMessageText, "🇷🇺 Русский", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		UserLanguage[update.Message.Chat.ID] = "ru"
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "🎉 Добро пожаловать в Monigo! Используйте клавиатуру ниже для доступа к командам:",
			ReplyMarkup: CreateMainKeyboardRU(),
		})
	})

	b.RegisterHandler(bot.HandlerTypeMessageText, "🇬🇧 English", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		UserLanguage[update.Message.Chat.ID] = "en"
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "🎉 Welcome to Monigo! Use the keyboard below to access commands:",
			ReplyMarkup: CreateMainKeyboardEN(),
		})
	})

	// Language change handler
	b.RegisterHandler(bot.HandlerTypeMessageText, "🔄 Сменить язык", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Выберите язык / Select language:",
			ReplyMarkup: CreateLanguageKeyboard(),
		})
	})

	b.RegisterHandler(bot.HandlerTypeMessageText, "🔄 Change language", bot.MatchTypeExact, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Выберите язык / Select language:",
			ReplyMarkup: CreateLanguageKeyboard(),
		})
	})
}
