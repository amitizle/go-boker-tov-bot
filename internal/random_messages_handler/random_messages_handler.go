package random_messages_handler

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"math/rand"
	"time"
)

var (
	replyProbability = 0.1
	replies          = []string{
		"מה הדיבור @%s?",
		"מה נז׳מע @%s",
		"איפה תלך @%s",
		"הכל רמייה @%s",
		"@%s שב בלול",
		"@%s על הצוואר שרשרת מגן דוד ומלשינים דוקר תמיד",
		"@%s אדידס לובש בנות כובש",
		"הרגת אותנו @%s",
		"מהנרגילה שואף ואת כל המלשינים מכפכף, לא ככה @%s?",
		"הרגת אותנו @%s",
	}
)

func Callback() func(*tgbotapi.BotAPI, *tgbotapi.Message) {
	return func(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
		if rand.Float64() <= replyProbability {
			response := replies[rand.Intn(len(replies))]
			msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(response, message.From))
			bot.Send(msg)
		}
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}
