package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/morheus9/go/src/main/find_data"
	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		var F = strconv.Itoa(find_data.Find_data())
		return c.Send(F)
	})

	b.Start()
}
