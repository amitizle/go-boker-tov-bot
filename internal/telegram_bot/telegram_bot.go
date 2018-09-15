package telegram_bot

import (
	"github.com/amitizle/go-boker-tov-bot/internal/simple_regex_handler"
	"github.com/amitizle/go-boker-tov-bot/internal/version_handler"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

var (
	bot    *tgbotapi.BotAPI
	router *Router
)

type BotConfig struct {
	Token          string
	WebhookAddress string
	Debug          bool
}

func New() *BotConfig {
	return &BotConfig{
		Debug: false,
	}
}

func Start(botConfig *BotConfig) {
	log.Println("Starting bot")
	bot, err := tgbotapi.NewBotAPI(botConfig.Token)
	if err != nil {
		log.Fatal(err)
	}
	router, err = NewRouter()
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = botConfig.Debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Handlers
	router.Handle("Version", version_handler.Callback())
	router.Handle("Simple Regex", simple_regex_handler.Callback())

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("Update Message: %v\n", update.Message)

		log.Printf("[From: %s] %s", update.Message.From.UserName, update.Message.Text)
		router.Route(bot, update.Message)
	}

}
