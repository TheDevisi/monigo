package telegram

import (
	"github.com/go-telegram/bot/models"
)

// Language selection keyboard
func CreateLanguageKeyboard() *models.ReplyKeyboardMarkup {
	return &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: "🇷🇺 Русский"},
				{Text: "🇬🇧 English"},
			},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}
}

// Main menu keyboard in Russian
func CreateMainKeyboardRU() *models.ReplyKeyboardMarkup {
	return &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: "📊 Статистика"}, // This will not be avalible in a few months
			},
			{
				{Text: "💻 CPU"},
				{Text: "💾 RAM"},
			},
			{
				{Text: "💿 Диск"},
				{Text: "⏱ Аптайм"},
			},
			{
				{Text: "🔄 Сменить язык"},
			},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
	}
}

// Main menu keyboard in English
func CreateMainKeyboardEN() *models.ReplyKeyboardMarkup {
	return &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: "📊 Statistics"}, // This will not be avalible in a few months
			},
			{
				{Text: "💻 CPU"},
				{Text: "💾 RAM"},
			},
			{
				{Text: "💿 Disk"},
				{Text: "⏱ Uptime"},
			},
			{
				{Text: "🔄 Change language"},
			},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
	}

}

func CreateStatKeyboard() *models.ReplyKeyboardMarkup {
	return &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: "123"},
			},
			{
				{Text: "321"},
			},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
	}
}
