package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nats-io/nats.go"
)

var (
	natsUrl             = os.Getenv("NATS_URL")
	telegramToken       = os.Getenv("TELEGRAM_BOT_TOKEN")
	chatId        int64 = 660704014 // Chat ID for my chat with the bot
)

func main() {
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("connecting to %s\n", natsUrl)
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err)
	}

	nc.QueueSubscribe("new_todo", "new_todo_queue", func(m *nats.Msg) {
		fmt.Printf("Received on [%s] message: %s\n", m.Subject, string(m.Data))

		msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("new todo\n```\n%s\n```", string(m.Data)))
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
	})
	nc.QueueSubscribe("updated_todo", "updated_todo_queue", func(m *nats.Msg) {
		fmt.Printf("Received on [%s] message: %s\n", m.Subject, string(m.Data))

		msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("updated todo\n```\n%s\n```", string(m.Data)))
		msg.ParseMode = "MarkdownV2"
		bot.Send(msg)
	})
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [new_todo]")

	runtime.Goexit()
}
