package main

import (
	"os"
	"github.com/op/go-logging"
	"github.com/kevinvandervlist/teshose/api"
	"github.com/kevinvandervlist/teshose/plugin"
	"github.com/kevinvandervlist/teshose/messages"
)

var log = logging.MustGetLogger("example")

func main() {
	key := os.Getenv("TELEGRAM_BOT_TOKEN")

	if key == "" {
		log.Info("Please provide a bot token in the environment variable 'TELEGRAM_BOT_API'.")
		os.Exit(0)
	}

	api := api.Create(log, key);
	api.Debug(true)
	err := api.Connect()

	commands := plugin.Create(log)
	broker := messages.Create(api)

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
			log.Info("Received a message from %s in %s(%d): %s", raw.Chat.FirstName, raw.Chat.Title, raw.Chat.ID ,raw.Text)
			go func() {
				inbound := broker.ConvertInbound(raw)

				resp, err := commands.Exec(raw.Text, &inbound)

				if(err != nil) {
					log.Error("An error occurred!", err)
				}

				if(resp.NoOp) {
					return
				}

				api.SendMessagesChannel <- broker.ConvertOutbound(resp)
			}()
		}
	}
}

