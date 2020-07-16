package tgbot

import (
	"bytes"
	"net/http"
	"shop/config"
	"shop/models"
	"text/template"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type ShopTgBot struct {
	tgBot  *tgbotapi.BotAPI
	chatId int64
	msgTemplateFile string
}

func NewShopTgBot(conf config.TelegramConfig) (*ShopTgBot, error) {
	cli := &http.Client{
		Timeout: 10 * time.Second,
	}
	bot, err := tgbotapi.NewBotAPIWithClient(conf.BotToken, cli)
	if err != nil {
		return nil, err
	}
	return &ShopTgBot{
		tgBot:           bot,
		chatId:          conf.ChatID,
		msgTemplateFile: conf.MsgTmplFile,
	}, nil
}

func (s *ShopTgBot) SendOrderNotification(order *models.Order) error {
	t, err := template.ParseFiles(s.msgTemplateFile)
	if err != nil {
		return err
	}

	var text bytes.Buffer
	if err := t.Execute(&text, order); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(s.chatId, text.String())

	_, err = s.tgBot.Send(msg)
	return err
}
