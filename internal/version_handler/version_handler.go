package version_handler

import (
	"gopkg.in/telegram-bot-api.v4"
	"regexp"
)

func Callback() func(*tgbotapi.BotAPI, *tgbotapi.Message) {
	return func(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
		r := regexp.MustCompile("/version")
		if r.MatchString(message.Text) {
			msg := tgbotapi.NewMessage(message.Chat.ID, "0.1.0")
			bot.Send(msg)
		}

	}
}

func init() {
	// Nothing to do
}
