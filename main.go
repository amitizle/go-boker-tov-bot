package main

import (
	bot "github.com/amitizle/go-boker-tov-bot/internal/telegram_bot"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var (
	mandatoryConfig = []string{
		"THE_CAT_API_KEY",
		"TELEGRAM_BOT_TOKEN",
		"TELEGRAM_BOT_WEBHOOK",
	}
	envVarPrefix = "BOKER_TOV"
)

func main() {
	configureBot()
	config := bot.New()
	config.TelegramBotToken = viper.GetString("TELEGRAM_BOT_TOKEN")
	config.WebhookAddress = viper.GetString("TELEGRAM_BOT_WEBHOOK")
	bot.Start(config)
}

func configureBot() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(envVarPrefix)
	missingConfigKeys := []string{}
	for _, k := range mandatoryConfig {
		if !viper.IsSet(k) {
			missingConfigKeys = append(missingConfigKeys, k)
		}
	}
	if len(missingConfigKeys) > 0 {
		log.Fatalf("Cannot start bot, missing %d mandatory config keys: %s",
			len(missingConfigKeys), strings.Join(missingConfigKeys, ", "))
	}
}
