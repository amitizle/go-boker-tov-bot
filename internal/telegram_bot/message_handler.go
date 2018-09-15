package telegram_bot

import (
	"gopkg.in/telegram-bot-api.v4"
	// "regexp"
)

type Handler struct {
	fun  func(*tgbotapi.BotAPI, *tgbotapi.Message)
	name string
}

type Router struct {
	handlers map[string]*Handler
}

func NewRouter() (*Router, error) {
	router := &Router{
		handlers: make(map[string]*Handler),
	}

	return router, nil
}

func (router *Router) Handle(name string, handlerFun func(*tgbotapi.BotAPI, *tgbotapi.Message)) error {
	// r := regexp.MustCompile(regex)
	router.handlers[name] = &Handler{
		fun:  handlerFun,
		name: name,
	}
	return nil // TODO return error
}

func (router *Router) Route(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	for _, handler := range router.handlers {
		go handler.fun(bot, message)
	}
}
