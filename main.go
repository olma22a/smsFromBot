package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"time"
)

var Bot *tgbotapi.BotAPI

func main() {
	bot, err := tgbotapi.NewBotAPI("7186314979:AAHVj1u0yjRKqGcH3qBtJVzJ4z3DlFjtqPw")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	m := map[int]string{
		1:  "хуец",
		2:  "залупец",
		3:  "тухлый",
		5:  "пенисистость",
		6:  "околовагиночки",
		7:  "нумизматище",
		8:  "скинхедТатар",
		9:  "Анальгетик",
		10: "блинница",
	}
	buttonText := "Матюк"
	button := tgbotapi.NewKeyboardButton(buttonText)

	keyboard := tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{button})
	message := tgbotapi.NewMessage(1398786974, "Матюк")
	message.ReplyMarkup = keyboard

	_, err = bot.Send(message)
	if err != nil {
		log.Panic(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			if update.Message.Text == "Матюк" {
				rand.Seed(time.Now().Unix())
				randomIndex := rand.Intn(len(m)) + 1
				randomWord := m[randomIndex]
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, randomWord))
			}
		}
	}

}
