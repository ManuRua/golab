package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
)

type application struct {
	client *tbot.Client
}

var (
	app   application
	bot   *tbot.Server
	token string
)

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	token = os.Getenv("TELEGRAM_TOKEN")
}

func main() {
	bot = tbot.New(token)
	app.client = bot.Client()
	bot.HandleMessage("/start", app.startHandler)
	log.Fatal(bot.Start())
}

// Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
	msg := "This is a bot whose sole purpose is to play rock, paper, scissors with you."
	a.client.SendMessage(m.Chat.ID, msg)
}
