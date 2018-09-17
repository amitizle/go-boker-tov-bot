package tch_handler

import (
	"gopkg.in/telegram-bot-api.v4"
	"regexp"
	"strings"
)

var (
	regex = regexp.MustCompile("ט+ח+")
)

func Callback() func(*tgbotapi.BotAPI, *tgbotapi.Message) {
	return func(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
		if regex.MatchString(message.Text) {
			tchToPff := regex.ReplaceAllStringFunc(message.Text, func(str string) string {
				hToP := strings.Replace(str, "ט", "פ", len(str))
				return strings.Replace(hToP, "ח", "ף", len(str))
			})
			msg := tgbotapi.NewMessage(message.Chat.ID, tchToPff)
			msg.ReplyToMessageID = message.MessageID
			bot.Send(msg)
		}
	}
}
