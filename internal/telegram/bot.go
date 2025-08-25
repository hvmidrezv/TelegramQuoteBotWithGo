package telegram

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"telegrambot/internal/config"
	"telegrambot/internal/quotes"
	"telegrambot/internal/scheduler"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tgbotapi.BotAPI

func Init(bot *tgbotapi.BotAPI) {
	Bot = bot
}

func HandleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}
	chatID := update.Message.Chat.ID
	userID := int64(update.Message.From.ID)

	if config.ConfigData.AdminID == 0 {
		config.ConfigData.AdminID = userID
		if err := config.AddGroup(chatID); err != nil {
			log.Println("AddGroup error:", err)
		}
		if err := config.Save(); err != nil {
			log.Println("Config save error:", err)
		}
		if err := scheduler.ScheduleFor(chatID, SendMessage); err != nil {
			log.Println("ScheduleFor error:", err)
		}
		if err := SendMessage(chatID, "✅ Admin & Group registered! فقط شما ادمین هستید."); err != nil {
			log.Println("SendMessage error:", err)
		}
		return
	}

	if !config.HasGroup(chatID) {
		if err := config.AddGroup(chatID); err != nil {
			log.Println("AddGroup error:", err)
		}
		if err := scheduler.ScheduleFor(chatID, SendMessage); err != nil {
			log.Println("ScheduleFor error:", err)
		}
		if err := SendMessage(chatID, "➕ گروه جدید ثبت شد. ارسال خودکار فعال شد."); err != nil {
			log.Println("SendMessage error:", err)
		}
	}

	if userID != config.ConfigData.AdminID {
		return
	}

	handleCommand(update.Message)
}

func handleCommand(msg *tgbotapi.Message) {
	text := strings.TrimSpace(msg.Text)
	switch {
	case strings.HasPrefix(text, "/setinterval"):
		parts := strings.Fields(text)
		if len(parts) != 2 {
			if err := SendMessage(msg.Chat.ID, "❌ فرمت دستور صحیح نیست. مثال: /setinterval 1"); err != nil {
				log.Println("SendMessage error:", err)
			}
			return
		}
		min, err := strconv.Atoi(parts[1])
		if err != nil || min <= 0 {
			if err := SendMessage(msg.Chat.ID, "❌ عدد دقیقه نامعتبر است"); err != nil {
				log.Println("SendMessage error:", err)
			}
			return
		}
		if err := config.SetIntervalFor(msg.Chat.ID, min); err != nil {
			log.Println("SetIntervalFor error:", err)
		}
		if err := scheduler.ScheduleFor(msg.Chat.ID, SendMessage); err != nil {
			log.Println("ScheduleFor error:", err)
		}
		if err := SendMessage(msg.Chat.ID, fmt.Sprintf("✔️ زمان ارسال پیام %d تنظیم شد", min)); err != nil {
			log.Println("SendMessage error:", err)
		}

	case text == "/now":
		q, err := quotes.FetchQuote()
		if err != nil {
			if err := SendMessage(msg.Chat.ID, "خطا در دریافت نقل‌قول"); err != nil {
				log.Println("SendMessage error:", err)
			}
			return
		}
		if err := SendMessage(msg.Chat.ID, quotes.FormatQuote(q)); err != nil {
			log.Println("SendMessage error:", err)
		}

	case text == "/groups":
		var b strings.Builder
		b.WriteString("📋 گروه‌های ثبت‌شده:\n")
		for _, gid := range config.ConfigData.Groups {
			b.WriteString(fmt.Sprintf("• %d (هر %d دقیقه)\n", gid, config.IntervalFor(gid)))
		}
		if err := SendMessage(msg.Chat.ID, b.String()); err != nil {
			log.Println("SendMessage error:", err)
		}
	}
}

func SendMessage(chatID int64, text string) error {
	// Replace <br> and <br/> with newlines
	text = strings.ReplaceAll(text, "<br>", "\n")
	text = strings.ReplaceAll(text, "<br/>", "\n")
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "HTML"
	_, err := Bot.Send(msg)
	if err != nil {
		log.Println("SendMessage error:", err)
	}
	return err
}
