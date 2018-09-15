package main

import (
	bot "github.com/amitizle/go-boker-tov-bot/internal/telegram_bot"
)

func main() {
	config := bot.New()
	config.Token = "" // TODO config
	bot.Start(config)
}
