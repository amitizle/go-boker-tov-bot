package telegram_bot

import (
	"github.com/amitizle/go-boker-tov-bot/internal/cat_handler"
	"github.com/amitizle/go-boker-tov-bot/internal/simple_regex_handler"
	"github.com/amitizle/go-boker-tov-bot/internal/version_handler"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	bot    *tgbotapi.BotAPI
	router *Router
)

type BotConfig struct {
	TelegramBotToken string
	WebhookAddress   string
	Debug            bool
}

func New() *BotConfig {
	return &BotConfig{
		Debug: false,
	}
}

func Start(botConfig *BotConfig) {
	log.Println("Starting bot")
	botAddr := "0.0.0.0:8686"
	log.Printf("Starting bot at %s", botAddr)
	bot, err := tgbotapi.NewBotAPI(botConfig.TelegramBotToken)
	if err != nil {
		log.Fatal(err)
	}
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(botConfig.WebhookAddress))
	if err != nil {
		log.Fatalf("Error while setting up webhook: %v", err)
	}
	webhookInfo, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatalf("Error while getting webhook info: %v", err)
	}
	if webhookInfo.LastErrorDate != 0 {
		log.Fatalf("[Telegram callback failed]%s", webhookInfo.LastErrorMessage)
	}

	router, err = NewRouter()
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = botConfig.Debug

	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(botAddr, nil)
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Handlers
	router.Handle("Version", version_handler.Callback())
	router.Handle("Simple Regex", simple_regex_handler.Callback())
	router.Handle("Cat", cat_handler.Callback())

	for update := range updates {
		if update.Message == nil {
			continue
		}
		router.Route(bot, update.Message)
	}

}

func init() {
	rand.Seed(time.Now().Unix())
}
