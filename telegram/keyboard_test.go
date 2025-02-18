package telegram

import (
	"github.com/go-telegram/bot/models"
)

// CreateMainKeyboard creates the main reply keyboard
func CreateMainKeyboard() *models.ReplyKeyboardMarkup {
	return &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: "/cpu"},
				{Text: "/ram"},
			},
			{
				{Text: "/disk"},
				{Text: "/uptime"},
			},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
	}
}

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
				{Text: "📊 Статистика"},
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
				{Text: "📊 Statistics"},
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
