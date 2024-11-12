package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tevian/internal/config"
)

type TelegramBot struct {
	bot *tgbotapi.BotAPI
}

func NewTGBot(cfg *config.Config) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.TGBot.Token)
	if err != nil {
		return nil, err
	}
	return &TelegramBot{bot: bot}, nil
}
