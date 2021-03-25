package mock_interface

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type TelegramMock interface {
	NewBotAPI(token string) (*tgbotapi.BotAPI, error)
	NewMessage(chatID int64, text string) tgbotapi.MessageConfig
	Send(tgbotapi.Chattable) (tgbotapi.MessageConfig, error)
}