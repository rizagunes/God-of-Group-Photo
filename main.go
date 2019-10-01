package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_SECRET_KEY"))
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.NewChatPhoto != nil && update.Message.From.UserName != bot.Self.UserName {
			photo := tgbotapi.NewSetChatPhotoUpload(update.Message.Chat.ID, os.Getenv("BOT_PHOTO"))
			bot.SetChatPhoto(photo)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, os.Getenv("BOT_MESSAGE"))
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

	}
}
