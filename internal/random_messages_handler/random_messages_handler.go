package random_messages_handler

import (
	"gopkg.in/telegram-bot-api.v4"
	"math/rand"
	"time"
)

var (
	replyProbability = 0.08
	replies          = []string{
		"מה נז׳מע",
		"מנשמע",
		"וואאההה",
		"פףף",
		"טחח",
		"איפה תלך",
		"הכל רמייה",
		"ניסים צ׳מע",
	}
)

func Callback() func(*tgbotapi.BotAPI, *tgbotapi.Message) {
	return func(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
		if rand.Float64() <= replyProbability {
			response := replies[rand.Intn(len(replies))]
			msg := tgbotapi.NewMessage(message.Chat.ID, response)
			bot.Send(msg)
		}
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}
