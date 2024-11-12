package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
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

func (b *TelegramBot) Start(ctx context.Context) {
	log.Printf("Started bot")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)
	for {
		select {
		case <-ctx.Done():
			return
		case update := <-updates:
			if update.Message != nil { // If we got a message
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID
				b.bot.Send(msg)
			}
		}
	}
}
