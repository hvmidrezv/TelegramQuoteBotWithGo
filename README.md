# 🧠 Telegram Quote Bot
![photo_2025-08-24_00-10-19](https://github.com/user-attachments/assets/a53c4626-1734-43ad-bf4b-a44da526c678)

A **Telegram bot** built with **Go** that sends **inspirational quotes** to Telegram groups at scheduled intervals. Ideal for community groups, productivity teams, or anyone who loves daily motivation.

---

## ✨ Features

- 🔁 **Scheduled Quotes**: Automatically delivers quotes to registered groups.
- 📚 **Powered by [ZenQuotes API](https://zenquotes.io/)**: Random inspirational quotes every time.
- 🛠️ **Admin Tools**:
    - `/setinterval <minutes>` – Set custom delivery intervals.
    - `/now` – Instantly send a quote.
- 🎨 **HTML-formatted Messages**: With time-based visual themes.

---

## 🚀 Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/yourusername/telegram-quote-bot.git
   cd telegram-quote-bot
   ```

2. **Install Dependencies**
   ```bash
   go mod tidy
   ```

3. **Setup Environment Variables**

   Create a `.env` file in the root directory:
   ```env
   BOT_TOKEN=your_telegram_bot_token
   ```

4. **Run the Bot**
   ```bash
   go run cmd/bot/main.go
   ```

---

## ⚙️ Commands

- `/setinterval <minutes>` – Set quote delivery interval for the current group.
- `/now` – Send a quote immediately.

---

## 🧩 Configuration

The bot uses a `config.json` file to persist:

- ✅ Admin user ID
- 🧑‍🤝‍🧑 Registered group IDs
- ⏱️ Delivery intervals

> This file is created and updated automatically as the bot runs.

---

## 📦 Dependencies

- [Go Telegram Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api)
- [GoCron](https://github.com/go-co-op/gocron)
- [godotenv](https://github.com/joho/godotenv)

---

## 🌐 API Used

- [ZenQuotes API](https://zenquotes.io/) – for motivational quote content.

---

## 📄 License

Licensed under the [MIT License](LICENSE).
