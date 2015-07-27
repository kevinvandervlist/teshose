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
		log.Info("Please provide a bot token in the environment variable 'TELEGRAM_BOT_TOKEN'.")
		os.Exit(0)
	}

	api := api.Create(log, key);
	err := api.Connect()
	api.Debug(true)

	commands := plugin.Create(log)

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
			// TODO: A pool of worker processes should handle this.
			go func() {
				pi := commands.BuildPluginInstance(raw.Text)
				pi.SetRequestMessage(raw)
				for ! pi.HasCompleted() {
					resp, err := pi.GetResponseMessage()

					if(err != nil) {
						log.Error("An error occurred!", err)
					}

					api.SendMessagesChannel <- resp
				}
			}()
		}
	}
}

