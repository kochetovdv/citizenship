package tgbot

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

/*
Получите токен от BotFather в Telegram.

Установите библиотеку для работы с Telegram API, например, "gopkg.in/telegram-bot-api.v4".

Этот код создает бота, который повторяет все сообщения, которые ему присылают.

*/

func main() {
	bot, err := tgbotapi.NewBotAPI("YOUR_TELEGRAM_BOT_TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
