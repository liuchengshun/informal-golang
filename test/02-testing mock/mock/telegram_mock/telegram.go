// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package provider

import (
	"strconv"

	tgbotapi "gopkg.in/telegram-bot-api.v4"

)

type TelegramMock interface {
	NewBotAPI(token string) (*tgbotapi.BotAPI, error)
	NewMessage(chatID int64, text string) tgbotapi.MessageConfig
	Send(tgbotapi.Chattable) (tgbotapi.MessageConfig, error)
}

type Message struct {
	Subj string
	To   []string
	From string
	Text string
}

type ProviderType struct {
	Value string
}

type TelegramConfig struct {
	Token string `yaml:"token"`
}

type Telegram struct {
	bot *tgbotapi.BotAPI
	ProviderType
}

func NewTelegram(config TelegramConfig) (*Telegram, error) {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		return nil, err
	}

	return &Telegram{
		bot,
		ProviderType{"robot"},
	}, nil
}

func (tg *Telegram) Send(message Message) error {
	for _, sChatID := range message.To {
		chatID, err := strconv.ParseInt(sChatID, 10, 64)
		if err != nil {
			return err
		}
		msg := tgbotapi.NewMessage(chatID, message.Text)
		_, err = tg.bot.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
