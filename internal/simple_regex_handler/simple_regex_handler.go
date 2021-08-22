package simple_regex_handler

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"math/rand"
	"regexp"
	"time"
)

type regexMatch struct {
	regex       *regexp.Regexp
	responses   []string
	probability float32
}

var (
	regexMatches = []*regexMatch{
		&regexMatch{
			probability: 0.5,
			regex:       wordBoundryRegexp("(שדמי|תום|נימר)"),
			responses:   []string{"אקסטרייייייייים", "תשמע הגזמנו אתמול", "וואה", "ווואייייי"},
		},
		&regexMatch{
			probability: 0.5,
			regex:       wordBoundryRegexp("(נימי|נמרוד|נימרוד|ניסים|נסים)"),
			responses:   []string{"פרושיאנטה", "רדיוהד", "סורי נרדמתי", "מת מת מת על טראמפ"},
		},
		&regexMatch{
			probability: 0.5,
			regex:       wordBoundryRegexp("(עמיתו|עמית|גולדברג|בנתן)"),
			responses:   []string{"¯\\_(ツ)_/¯", "רדיוהד מופעים פה מחר", "מפתיע", "יפתיע אותך?"},
		},
		&regexMatch{
			probability: 0.5,
			regex:       wordBoundryRegexp("(ירון|חתול|חאתול)"),
			responses:   []string{"טחח", "יא ז׳וז׳ו טחח", "פףף", "מה נז׳מע", "חתולה"},
		},
		&regexMatch{
			probability: 0.5,
			regex:       wordBoundryRegexp("(יואב|יוחאפי|יוחאפ)"),
      responses:   []string{"איפה המברשת שיניים שלי?", "הרוקו אחי", "נביא קצת טחינה עם כרובית?"},
		},
		&regexMatch{
			probability: 1.0,
			regex:       wordBoundryRegexp("(תומר פישמן|ניב מג'ר|ניב מג׳ר)"),
			responses:   []string{"מטונף"},
		},
	}
)

func Callback() func(*tgbotapi.BotAPI, *tgbotapi.Message) {
	return func(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
		for _, match := range regexMatches {
			if match.regex.MatchString(message.Text) && rand.Float32() <= match.probability {
				response := match.responses[rand.Intn(len(match.responses))]
				msg := tgbotapi.NewMessage(message.Chat.ID, response)
				bot.Send(msg)
			}
		}
	}
}

func init() {
	rand.Seed(time.Now().Unix())
	// Nothing to do
}

func wordBoundryRegexp(words string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf("(?:\\A|\\s)%s(?:\\s|\\z)", words))
}
