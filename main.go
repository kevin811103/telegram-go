package main

import (
    "log"
    "os"

    tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
    var (
        port      = os.Getenv("PORT")
        publicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
        token     = os.Getenv("TOKEN")      // you must add it to your config vars
    )

    webhook := &tb.Webhook{
        Listen:   ":" + port,
        Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
    }

    pref := tb.Settings{
        Token:  token,
        Poller: webhook,
    }

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hi!")
	})
	
    b, err := tb.NewBot(pref)
    if err != nil {
        log.Fatal(err)
    }
}