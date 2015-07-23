package main

import (
	"os"
	"github.com/op/go-logging"
	"github.com/kevinvandervlist/teshose/api"
	"github.com/kevinvandervlist/teshose/plugin"
)

var log = logging.MustGetLogger("example")

func main() {
	key := os.Getenv("TELEGRAM_BOT_TOKEN")

	if key == "" {
		log.Info("Please provide a bot token in the environment variable 'TELEGRAM_BOT_API'.")
		os.Exit(0)
	}

	commands := plugin.Create()

	/*
	foo, _ := plugin.Exec("Foo")
	bar, _ := plugin.Exec("Bar")
	baz, _ := plugin.Exec("Baz")

	fmt.Printf("Foo: %s\n", foo)
	fmt.Printf("Bar: %s\n", bar)
	fmt.Printf("Bar: %s\n", baz)
	*/

	api := api.Create(log, key);
	api.Debug(true)
	err := api.Connect()

	if(err != nil) {
		log.Critical("A connection error occurred: ", err)
		panic(err)
	}

	for {
		select {
		case e := <- api.ErrorChannel:
			log.Error("An error occurred: ", e)
			panic(e)
			return
		case raw := <- api.ReceiveMessagesChannel:
			log.Info("A message is received: %s", raw.Text)
			//cmd := commands.parse(raw.Text)
			m := plugin.IncomingMessage{
				Text: raw.Text,
			}
			func() {
				resp := commands.Exec("echo", m)
				api.SendMessagesChannel <- actualResponse
			}()
		}
	}
}

