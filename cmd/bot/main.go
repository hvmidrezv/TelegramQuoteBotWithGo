package main

import (
	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"telegrambot/internal/config"
	"telegrambot/internal/scheduler"
	"telegrambot/internal/telegram"
	"time"
)

func main() {
	_ = godotenv.Load()

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN not set")
	}

	err := config.Load()
	if err != nil {
		log.Fatal("خطا در لود config:", err)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	telegram.Init(bot)

	s := gocron.NewScheduler(time.UTC)
	scheduler.Start(s, telegram.SendMessage)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		telegram.HandleUpdate(update)
	}
}
