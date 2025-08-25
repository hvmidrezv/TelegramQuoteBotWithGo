# Telegram Quote Bot 🇬🇧➡️🇮🇷

A Telegram bot written in Golang that:
- Fetches English quotes from [ZenQuotes.io](https://zenquotes.io/keywords/love)
- Translates them to Persian using [LibreTranslate](https://libretranslate.com)
- Sends them automatically in your group
- Admin and group auto-detected at first use
- Admin can control interval via commands

## 🔧 Setup

1. Copy `.env.example` to `.env` and set your Telegram bot token.
2. Run the bot:

```bash
go run main.go
```

## 🛠️ Commands

- `/setinterval <minutes>` – Set the interval between quotes (admin only)
- `/now` – Send a quote immediately (admin only)
